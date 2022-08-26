import React, { useEffect } from 'react';
import { Route, Routes } from 'react-router-dom';
import Home from './pages/Home';
import About from './pages/About';
import { DownloadInfo, syncDownloads } from './store/downloadsSlice';
import { eventsEmit, eventsOff, eventsOn } from './binds/runtime';
import { EVENT_NAME_BACKEND_STATE_CHANGED, EVENT_NAME_FRONTEND_READY } from './constants/events';
import { BackendState } from './types';
import { useDispatch } from 'react-redux';
import { AppDispatch } from './store';

function App() {
  const dispatch = useDispatch<AppDispatch>();

  useEffect(() => {
    eventsOn(EVENT_NAME_BACKEND_STATE_CHANGED, (data: object) => {
      const backendState = data as BackendState;
      dispatch(syncDownloads(backendState.downloads));
    });

    eventsEmit(EVENT_NAME_FRONTEND_READY, {});

    return () => {
      eventsOff(EVENT_NAME_BACKEND_STATE_CHANGED);
    };
  }, []);

  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/about" element={<About />} />
    </Routes>
  );
}

export default App;
