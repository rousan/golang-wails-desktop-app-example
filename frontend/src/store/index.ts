import { configureStore } from '@reduxjs/toolkit';
import downloadsSlice from './downloadsSlice';

const store = configureStore({
  reducer: {
    downloads: downloadsSlice
  },
});

export type AppState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch

export default store;