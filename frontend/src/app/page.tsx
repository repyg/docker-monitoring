"use client";

import { Spin } from "antd";
import ContainerTable from "./components/ContainerTable";
import ContainerStatusChart from "./components/ContainerStatusChart";
import styles from "./page.module.css";
import { useQuery } from "@tanstack/react-query";
import { fetchContainers, Container } from "./api/containers";


export default function Home() {
  const { data: containers, error, isLoading } = useQuery<Container[]>({
    queryKey: ["containers"],
    queryFn: fetchContainers,
    refetchInterval: 5000,
  });

  if (isLoading) {
    return (
      <div style={{ display: "flex", justifyContent: "center", marginTop: 20 }}>
        <Spin tip="Загрузка контейнеров..." size="large" />
      </div>
    );
  }

  if (error) {
    return <div>Ошибка загрузки данных</div>;
  }
  return (
    <div className={styles.pageWrapper}>
      <header className={styles.header}>
        <h1 className={styles.headerTitle}>Docker Monitoring App</h1>
      </header>

      <section className={styles.tableSection}>
        <ContainerTable />
        {containers && containers.length > 0 && (
          <ContainerStatusChart containers={containers} />
        )}
      </section>

      <footer className={styles.footer}>
        made by <strong>repyg</strong>
      </footer>
    </div>
  );
}