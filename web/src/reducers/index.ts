import { combineReducers } from "@reduxjs/toolkit";

import { AuthState, authReducer } from "./authReducer";

interface RootState {
  auth: AuthState;
}

export const rootReducer = combineReducers<RootState | undefined>({
  auth: authReducer,
});
