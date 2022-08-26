import React, { useEffect } from "react";
import { Link } from "react-router-dom";
import styles from './index.module.css';

export default function About() {
  useEffect(() => {
    console.log("About is mounted");
    return () => {
      console.log("About is unmounted");
    };
  }, []);

  return (
    <div className={styles.about}>
      <div className={styles.label}>
        About
      </div>
      <Link to="/">Home</Link>
    </div>
  );
}