"use client";

import { useMemo } from "react";
import type { BreakdownItem } from "@/lib/validators/analytics";

const CHART_COLORS = [
  "oklch(0.68 0.17 55)", // primary orange
  "oklch(0.6 0.118 185)", // teal
  "oklch(0.75 0.16 70)", // gold
  "oklch(0.646 0.222 41)", // coral
  "oklch(0.769 0.188 70)", // amber
  "oklch(0.55 0.15 280)", // purple
  "oklch(0.6 0.15 150)", // green
  "oklch(0.5 0.1 220)", // blue
];

interface DonutChartProps {
  data: BreakdownItem[];
  size?: number;
  strokeWidth?: number;
}

export function DonutChart({
  data,
  size = 180,
  strokeWidth = 28,
}: DonutChartProps) {
  const segments = useMemo(() => {
    const total = data.reduce((sum, item) => sum + item.count, 0);
    if (total === 0) return [];

    const radius = (size - strokeWidth) / 2;
    const circumference = 2 * Math.PI * radius;
    let offset = 0;

    return data.map((item, i) => {
      const fraction = item.count / total;
      const length = fraction * circumference;
      const gap = 3; // gap between segments
      const seg = {
        key: item.key,
        count: item.count,
        percentage: Math.round(fraction * 100),
        color: CHART_COLORS[i % CHART_COLORS.length],
        dashArray: `${Math.max(length - gap, 0)} ${circumference - Math.max(length - gap, 0)}`,
        dashOffset: -offset,
        radius,
      };
      offset += length;
      return seg;
    });
  }, [data, size, strokeWidth]);

  const total = data.reduce((sum, item) => sum + item.count, 0);

  if (data.length === 0) {
    return (
      <div
        className="flex items-center justify-center text-muted-foreground text-sm"
        style={{ width: size, height: size }}
      >
        No data
      </div>
    );
  }

  return (
    <div className="flex flex-col sm:flex-row items-center gap-6">
      {/* Donut */}
      <div className="relative shrink-0" style={{ width: size, height: size }}>
        <svg
          viewBox={`0 0 ${size} ${size}`}
          className="w-full h-full -rotate-90"
        >
          {segments.map((seg) => (
            <circle
              key={seg.key}
              cx={size / 2}
              cy={size / 2}
              r={seg.radius}
              fill="none"
              stroke={seg.color}
              strokeWidth={strokeWidth}
              strokeDasharray={seg.dashArray}
              strokeDashoffset={seg.dashOffset}
              strokeLinecap="round"
              className="transition-all duration-500"
            />
          ))}
        </svg>
        {/* Center text */}
        <div className="absolute inset-0 flex flex-col items-center justify-center">
          <p className="text-2xl font-bold">{total.toLocaleString()}</p>
          <p className="text-xs text-muted-foreground">Total</p>
        </div>
      </div>

      {/* Legend */}
      <div className="flex flex-col gap-2 min-w-0">
        {segments.map((seg) => (
          <div key={seg.key} className="flex items-center gap-2.5 text-sm">
            <div
              className="h-3 w-3 rounded-full shrink-0"
              style={{ backgroundColor: seg.color }}
            />
            <span className="truncate text-muted-foreground">{seg.key}</span>
            <span className="ml-auto font-semibold tabular-nums shrink-0">
              {seg.percentage}%
            </span>
          </div>
        ))}
      </div>
    </div>
  );
}
