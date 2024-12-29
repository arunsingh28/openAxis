import { createSlice } from '@reduxjs/toolkit';

import { IController } from '@/types/controller';

const initialState: IController[] = [];

const apiCenterSlice = createSlice({
  name: 'apiCenter',
  initialState,
  reducers: {
    newController: (_, action) => {
      return action.payload;
    }
  }
});

export const { newController } = apiCenterSlice.actions;
export default apiCenterSlice.reducer;
