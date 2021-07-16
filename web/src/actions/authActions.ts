import { createAction } from "@reduxjs/toolkit";

export const SET_TOKEN = "SET_TOKEN";

export const setToken = createAction(SET_TOKEN, (token) => {
  return {
    payload: {
      token: token,
    },
  };
});
