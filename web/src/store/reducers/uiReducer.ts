import { createReducer } from "@reduxjs/toolkit";

import { setScrollable } from "../actions/uiActions";

export interface UIState {
  overlayState: boolean;
}

const defaultState: UIState = {
  overlayState: false,
};

export const uiReducer = createReducer(defaultState, (builder) => {
  builder.addCase(setScrollable, (state, action) => {
    state.overlayState = action.payload;
  });
});
