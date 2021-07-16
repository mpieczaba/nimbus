import React from "react";
import { Formik, FormikErrors } from "formik";
import { IconUser, IconLock } from "@tabler/icons";

import { useAppDispatch } from "../../hooks/store";

import { setToken } from "../../actions/authActions";

import {
  LoginMutationVariables,
  useLoginMutation,
} from "../../generated/graphql";

import Form from "../Form";

import { Wrapper, Header, Error } from "./styles";

import Input, {
  InputWrapper,
  InputAndIconWrapper,
  InputIcon,
  InputError,
} from "../Input";
import Button from "../Button";

const LoginForm: React.FC = () => {
  const dispatch = useAppDispatch();

  const [login, { data, error }] = useLoginMutation();

  return (
    <Wrapper>
      <Formik
        initialValues={{
          username: "",
          password: "",
        }}
        validate={({ username, password }: LoginMutationVariables) => {
          const errors: FormikErrors<LoginMutationVariables> = {};

          if (!username) errors.username = "Required!";
          else if (!password) errors.password = "Required!";
          else if (!/^(?:[A-Za-z0-9_]+)$/.test(username))
            errors.username = "Invalid username!";

          return errors;
        }}
        onSubmit={async (values: LoginMutationVariables, { setSubmitting }) => {
          console.log(values);

          try {
            await login({ variables: values });

            dispatch(setToken(data?.login.token));
          } catch (err) {
            console.log(err);
          }

          setSubmitting(false);
        }}
      >
        {({
          isSubmitting,
          handleReset,
          handleSubmit,
          handleChange,
          errors,
        }) => (
          <Form onReset={handleReset} onSubmit={handleSubmit}>
            <Header>Welcome back!</Header>

            <InputWrapper>
              <InputAndIconWrapper>
                <InputIcon>
                  <IconUser />
                </InputIcon>

                <Input
                  type="text"
                  name="username"
                  placeholder="Username"
                  onChange={handleChange}
                />
              </InputAndIconWrapper>

              <InputError>{errors.username}</InputError>
            </InputWrapper>

            <InputWrapper>
              <InputAndIconWrapper>
                <InputIcon>
                  <IconLock />
                </InputIcon>

                <Input
                  type="password"
                  name="password"
                  placeholder="Password"
                  onChange={handleChange}
                />
              </InputAndIconWrapper>

              <InputError>{errors.password}</InputError>
            </InputWrapper>

            {error && <Error>{error.message}</Error>}

            <Button type="submit" disabled={isSubmitting}>
              Sign in!
            </Button>
          </Form>
        )}
      </Formik>
    </Wrapper>
  );
};

export default LoginForm;
