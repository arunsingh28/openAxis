import { configureStore } from "@reduxjs/toolkit";

import projectDeducer from "./slices/projectSlice";

export const store = configureStore({
  reducer: {
    project: projectDeducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
