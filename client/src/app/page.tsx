"use client";
import styles from "./page.module.css";
import { Suspense } from "react";
import Post from "@/app/post";

export default function Home() {
  return (
    <main className={styles.main}>
      <Suspense fallback={<div>Loading</div>}>
        <Post/>
      </Suspense>
    </main>
  );
}
