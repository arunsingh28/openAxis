import { createSlice } from '@reduxjs/toolkit';

import { IProject } from '@/types/project';

const initialState: IProject[] = [];

export const projectSlice = createSlice({
  name: 'project',
  initialState,
  reducers: {
    initProject: (_, action) => {
      return action.payload;
    }
  }
});

export const { initProject } = projectSlice.actions;
export default projectSlice.reducer;
