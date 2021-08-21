import React, { ChangeEvent, useState } from "react";
import { Formik, Form, FormikErrors } from "formik";
import { IconPencil, IconPlus } from "@tabler/icons";

import {
  FilesQuery,
  FilesDocument,
  useCreateFileMutation,
  FilesQueryVariables,
  FileEdge,
} from "../../generated/graphql";

import Button from "../Button";
import Popup, { PopupItem } from "../Popup";
import Input, {
  InputWrapper,
  InputAndIconWrapper,
  InputIcon,
  InputError,
} from "../Input";

import { Wrapper, FileInputLabel, PopupItemName } from "./styles";

const UploadButton: React.FC = () => {
  const [createFile, { error }] = useCreateFileMutation({
    onCompleted: () => showHidePopup(false),
    update: (cache, { data }) => {
      if (data?.createFile) {
        const existingFiles = cache.readQuery<FilesQuery, FilesQueryVariables>({
          query: FilesDocument,
          variables: { name: null, tags: [] },
        });

        cache.writeQuery<FilesQuery>({
          query: FilesDocument,
          variables: { name: null, tags: [] },
          data: {
            files: {
              edges: [
                { node: data.createFile, cursor: "" },
                ...(existingFiles!.files!.edges as Array<FileEdge>),
              ],
              pageInfo: existingFiles!.files!.pageInfo,
            },
          },
        });
      }
    },
  });

  const [popup, showHidePopup] = useState<boolean>(false);

  return (
    <Wrapper>
      <Formik
        initialValues={{
          upload: File,
          name: "",
        }}
        validate={({ name }) => {
          const errors: FormikErrors<{ name: string }> = {};

          if (!/^[^/>|:&]+$/.test(name)) errors.name = "Invalid file name!";

          return errors;
        }}
        onSubmit={async (values, { setSubmitting }) => {
          try {
            await createFile({
              variables: {
                fileInput: {
                  file: values.upload,
                  name: values.name,
                },
              },
            });
          } catch (err) {
            console.log(err);
          }

          setSubmitting(false);
        }}
      >
        {({ isSubmitting, setFieldValue, values, handleChange, errors }) => (
          <Form>
            <FileInputLabel>
              <input
                type="file"
                name="upload"
                hidden
                onChange={(e: ChangeEvent<HTMLInputElement>) => {
                  if (e.target.files) {
                    setFieldValue("upload", e.target.files[0]);
                    setFieldValue("name", e.target.files[0].name);

                    showHidePopup(true);
                  }
                }}
              />

              <IconPlus size="2rem" />
            </FileInputLabel>

            <Popup active={popup} hidePopup={() => showHidePopup(false)}>
              <PopupItemName>Change file name</PopupItemName>

              <PopupItem>
                <InputWrapper>
                  <InputAndIconWrapper>
                    <InputIcon>
                      <IconPencil />
                    </InputIcon>

                    <Input
                      type="text"
                      name="name"
                      placeholder="File name"
                      value={values.name}
                      onChange={handleChange}
                    />
                  </InputAndIconWrapper>

                  <InputError>{errors.name}</InputError>
                </InputWrapper>
              </PopupItem>

              <PopupItem right>
                <Button type="submit" disabled={isSubmitting}>
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
