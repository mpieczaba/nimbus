FROM golang:1.15-alpine as builder

RUN apk add -U build-base sudo git curl cmake make libc-dev linux-headers pkgconfig ffmpeg-dev gstreamer-dev

ENV PKG_CONFIG_PATH=/usr/local/lib64/pkgconfig
ENV CGO_CXXFLAGS="--std=c++1z"
ENV CGO_CPPFLAGS=-I/usr/local/include
ENV LD_LIBRARY_PATH=/usr/local/lib64
ENV CGO_LDFLAGS="-L/usr/local/lib -lopencv_core -lopencv_face -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_objdetect -lopencv_features2d -lopencv_video -lopencv_dnn -lopencv_xfeatures2d -lopencv_plot -lopencv_tracking"

RUN GO111MODULE=off go get -u -d gocv.io/x/gocv && cd $GOPATH/src/gocv.io/x/gocv && \
    # make install https://github.com/hybridgroup/gocv/issues/819
    make deps && make download && make sudo_pre_install_clean && \
    # make build && \
    cd /tmp/opencv/opencv-4.5.1 && \
    mkdir build && \
    cd build && \
    rm -rf * && \
    cmake -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/usr/local -D BUILD_SHARED_LIBS=ON \
    	 -D OPENCV_EXTRA_MODULES_PATH=/tmp/opencv/opencv_contrib-4.5.1/modules -D BUILD_DOCS=OFF -D \
    	  BUILD_EXAMPLES=OFF -D BUILD_TESTS=OFF -D BUILD_PERF_TESTS=OFF -D BUILD_opencv_java=NO -D \
    	  BUILD_opencv_python=NO -D BUILD_opencv_python2=NO -D BUILD_opencv_python3=NO -D WITH_JASPER=OFF \
    	  -D OPENCV_GENERATE_PKGCONFIG=ON .. && \
    make -j $(nproc --all) && \
    make preinstall && \
    cd $GOPATH/src/gocv.io/x/gocv && \
    make sudo_install && make clean && make verify

RUN pkg-config --cflags --libs opencv4

WORKDIR /src

COPY . .

RUN go build -o nimbus .

FROM alpine

RUN apk add -U libstdc++ ffmpeg gstreamer

COPY --from=builder /usr/local/lib64 /usr/local/lib64
COPY --from=builder /usr/local/lib64/pkgconfig/opencv4.pc /usr/local/lib64/pkgconfig/opencv4.pc
COPY --from=builder /usr/local/include/opencv4 /usr/local/include/opencv4

WORKDIR /app

COPY --from=builder /src/nimbus .

ENV PKG_CONFIG_PATH=/usr/local/lib64/pkgconfig
ENV CGO_CXXFLAGS="--std=c++1z"
ENV CGO_CPPFLAGS=-I/usr/local/include
ENV LD_LIBRARY_PATH=/usr/local/lib64
ENV CGO_LDFLAGS="-L/usr/local/lib -lopencv_core -lopencv_face -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_objdetect -lopencv_features2d -lopencv_video -lopencv_dnn -lopencv_xfeatures2d -lopencv_plot -lopencv_tracking"

CMD ["/app/nimbus", "start"]
