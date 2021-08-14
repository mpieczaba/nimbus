import { createAction } from "@reduxjs/toolkit";

const SHOW_HIDE_OVERLAY = "SHOW_HIDE_OVERLAY";

export const setScrollable = createAction(SHOW_HIDE_OVERLAY, (scrollState) => {
  return {
    payload: scrollState,
  };
});
