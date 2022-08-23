import React, { useEffect } from "react";
import './index.css';
import { Link } from 'react-router-dom';

function HomePage() {
  useEffect(() => {
    console.log("Home is mounted");
    return () => {
      console.log("Home is unmounted");
    };
  }, []);

  return (
    <div className="home">
      Home
      <Link to="/about">About</Link>
    </div>
  );
}

export default HomePage;