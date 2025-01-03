import { configureStore } from '@reduxjs/toolkit';
import shifuReducer from './shifuSlice';

export const store = configureStore({
  reducer: {
    shifu: shifuReducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch; 