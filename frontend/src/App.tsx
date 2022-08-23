import React, { useEffect } from 'react';
import { Route, Routes } from 'react-router-dom';
import AboutPage from './pages/About';
import HomePage from './pages/Home';

function App() {
  useEffect(() => {
    console.log("App is mounted");
    return () => {
      console.log("App is unmounted");
    };
  }, []);

  return (
    <Routes>
      <Route path="/" element={<HomePage />} />
      <Route path="/about" element={<AboutPage />} />
    </Routes>
  );
}

export default App;
