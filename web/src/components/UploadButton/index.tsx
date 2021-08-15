import React, { ChangeEvent, useState } from "react";
import { extname } from "path";
import { Formik, Form } from "formik";
import { IconEdit, IconPlus, IconTag } from "@tabler/icons";

import { useCreateFileMutation } from "../../generated/graphql";

import Button from "../Button";
import Popup, { PopupItem, PopupItemIcon } from "../Popup";
import FileThumbnail from "../FileThumbnail";

import { Wrapper, FileInputLabel, Thumbnail, PopupItemName } from "./styles";

const UploadButton: React.FC = () => {
  const [createFile, { error }] = useCreateFileMutation({
    onCompleted: ({ createFile }) => {
      if (createFile) console.log(createFile);
    },
  });

  const [popup, showHidePopup] = useState<boolean>(false);

  return (
    <Wrapper>
      <Formik
        initialValues={{
          upload: File,
        }}
        onSubmit={async (values, { setSubmitting }) => {
          console.log(values.upload);

          try {
            await createFile({
              variables: {
                fileInput: {
                  file: values.upload,
                },
              },
            });
          } catch (err) {
            console.log(err);
          }

          setSubmitting(false);
        }}
      >
        {({ isSubmitting, submitForm, setFieldValue, values }) => (
          <Form>
            <FileInputLabel>
              <input
                type="file"
                name="upload"
                hidden
                onChange={(e: ChangeEvent<HTMLInputElement>) => {
                  if (e.target.files) {
                    setFieldValue("upload", e.target.files[0]);

                    showHidePopup(true);
                  }
                }}
              />

              <IconPlus size="2rem" />
            </FileInputLabel>

            <Popup active={popup} hidePopup={() => showHidePopup(false)}>
              <PopupItemName>
                <Thumbnail>
                  {<FileThumbnail extension={extname(values.upload.name)} />}
                </Thumbnail>
                <span>{values.upload.name}</span>
              </PopupItemName>

              <PopupItem>
                <PopupItemIcon>
                  <IconEdit />
                </PopupItemIcon>
                Rename
              </PopupItem>

              <PopupItem>
                <PopupItemIcon>
                  <IconTag />
                </PopupItemIcon>
                Add tags
              </PopupItem>

              <PopupItem>
                <Button
                  type="submit"
                  disabled={isSubmitting}
                  onClick={async (e) => {
                    await submitForm();
                  }}
                >
                  Upload
                </Button>
                {error && console.log(error.message)}
              </PopupItem>
            </Popup>
          </Form>
        )}
      </Formik>
    </Wrapper>
  );
};

export default UploadButton;
