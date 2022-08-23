import React, { useEffect } from "react";
import { Link } from "react-router-dom";
import './index.css';

function AboutPage() {
  useEffect(() => {
    console.log("About is mounted");
    return () => {
      console.log("About is unmounted");
    };
  }, []);


  return (
    <div className="about">
      About
      <Link to="/">Home</Link>
    </div>
  );
}

export default AboutPage;