# Nimbus ☁️
Nimbus - extensible storage system focused on quick data access

### Before you start read the disclaimer (at the end of this file)!

## 🚀 Running the app
1. Make sure that you have installed [Go](https://golang.org/), [GoCV](https://gocv.io/) and [MySQL](https://www.mysql.com/).
   Do not forget to create a database! 
2. Clone this repository to your local machine:
```
$ git clone https://github.com/mpieczaba/nimbus && cd nimbus
```
3. Create .env file. Sample content can be found in the `.env.example` file.
4. Run the app!
```
$ go run . start
```

## ❓ How does it work?
All cloud-based storage systems use slow and boring hierarchical directory systems.
Using them, it takes ages to find the file you are looking for.

Our solution for this problem is quite unusual - what if all the files were stored in one directory?

Files can be grouped into multiple tags by their content. 
Furthermore, algorithms can classify a file by its name and content, using such things as computer vision, reverse image search and more.

Combined with a robust search engine, Nimbus delivers your files **in the blink of an eye**!

## 🔥 Features

- Easy to use and (wannabe) safe GraphQL API
- ~~Web service and dashboard~~
- ~~Event hooks~~
- ~~Computer vision classifiers~~
- ~~CLI management tools~~


## ⚠️ Disclaimer
Nimbus is still in development! Use it at your own risk!

We do not guarantee that the code is bug-free (it probably has plenty of them).

Some features are truncated to provide at least a working prototype.

## 👏 Contributing
All pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.