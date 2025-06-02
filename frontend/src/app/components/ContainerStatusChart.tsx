// src/app/components/ContainerStatusChart.tsx
"use client";

import { BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer, Cell } from "recharts";
import { Container } from "../api/containers";
import { useMemo } from "react";
import styles from "../page.module.css";

interface ContainerStatusChartProps {
  containers: Container[];
}

const statusColors: { [key: string]: string } = {
  running: "#36B37E",
  paused: "#0052CC",
  exited: "#FF5630",
  unknown: "#FFAB00",
};

export default function ContainerStatusChart({ containers }: ContainerStatusChartProps) {
  const statusCounts = useMemo(() => {
    const counts: { [key: string]: number } = {};
    containers.forEach((container) => {
      counts[container.status] = (counts[container.status] || 0) + 1;
    });
    return Object.keys(counts).map((status) => ({
      status: status.charAt(0).toUpperCase() + status.slice(1),
      count: counts[status],
      color: statusColors[status] || statusColors.unknown,
    }));
  }, [containers]);

  return (
    <div style={{ width: "100%", height: 300, marginTop: 40 }}>
      {/* Применяем классы header и headerTitle для стилизации заголовка */}
      <div className={styles.header} style={{ padding: "10px 0", marginBottom: "20px" }}>
        <h3 className={styles.headerTitle} style={{ margin: 0 }}>
          Распределение статусов контейнеров
        </h3>
      </div>
      <ResponsiveContainer>
        <BarChart
          data={statusCounts}
          margin={{
            top: 20,
            right: 30,
            left: 20,
            bottom: 5,
          }}
        >
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="status" />
          <YAxis allowDecimals={false} />
          <Tooltip />
          <Legend />
          <Bar dataKey="count" name="Количество контейнеров">
            {
              statusCounts.map((entry, index) => (
                <Cell key={`cell-${index}`} fill={entry.color} />
              ))
            }
          </Bar>
        </BarChart>
      </ResponsiveContainer>
    </div>
  );
}