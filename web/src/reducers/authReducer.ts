import { createReducer } from "@reduxjs/toolkit";

import { setToken } from "../actions/authActions";

export interface AuthState {
  readonly token?: string | null;
}

const defaultState: AuthState = {
  token: localStorage.getItem("token"),
};

export const authReducer = createReducer(defaultState, (builder) => {
  builder.addCase(setToken, (state, action) => {
    localStorage.setItem("token", action.payload.token);
  });
});
