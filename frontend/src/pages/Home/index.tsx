import React, { useEffect } from "react";
import { Link } from 'react-router-dom';
import styles from './index.module.css';

function Home() {
  useEffect(() => {
    console.log("Home is mounted");
    return () => {
      console.log("Home is unmounted");
    };
  }, []);

  return (
    <div className={styles.home}>
      <div className={styles.label}>
        Home
      </div>
      <Link to="/about">About</Link>
    </div>
  );
}

export default Home;