import { combineReducers } from "@reduxjs/toolkit";

import { AuthState, authReducer } from "./authReducer";
import { uiReducer, UIState } from "./uiReducer";

interface RootState {
  auth: AuthState;
  ui: UIState;
}

export const rootReducer = combineReducers<RootState | undefined>({
  auth: authReducer,
  ui: uiReducer,
});
