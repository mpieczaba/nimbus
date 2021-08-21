import React, { useEffect } from "react";
import { useHistory } from "react-router-dom";
import { Formik, FormikErrors } from "formik";
import { IconUser, IconLock } from "@tabler/icons";

import { useAppDispatch } from "../../hooks/store";

import { setToken } from "../../store/actions/authActions";

import {
  LoginMutationVariables,
  useLoginMutation,
} from "../../generated/graphql";

import Form, { Header } from "../Form";

import { Wrapper, Logo, Error } from "./styles";

import Input, {
  InputWrapper,
  InputAndIconWrapper,
  InputIcon,
  InputError,
} from "../Input";
import Button from "../Button";

const LoginForm: React.FC = () => {
  const dispatch = useAppDispatch();
  const history = useHistory();

  const [login, { error }] = useLoginMutation({
    onCompleted: ({ login }) => {
      if (login) {
        dispatch(setToken(login.token));

        history.push("/");
      }
    },
  });

  useEffect(() => {
    if (localStorage.getItem("token")) history.push("/");
  }, [history]);

  return (
    <Wrapper>
      <Logo>Nimbus</Logo>

      <Formik
        initialValues={{
          username: "",
          password: "",
        }}
        validate={({ username, password }: LoginMutationVariables) => {
          const errors: FormikErrors<LoginMutationVariables> = {};

          if (!username) errors.username = "Required!";
          else if (!password) errors.password = "Required!";
          else if (!/^:[A-Za-z0-9_]+$/.test(username))
            errors.username = "Invalid username!";

          return errors;
        }}
        onSubmit={async (values: LoginMutationVariables, { setSubmitting }) => {
          try {
            await login({
              variables: values,
            });
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
