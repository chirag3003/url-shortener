"use client";

import { useMemo, useState } from "react";
import type { AnalyticsPoint } from "@/lib/validators/analytics";

interface AreaChartProps {
  data: AnalyticsPoint[];
  height?: number;
  color?: string;
  showLabels?: boolean;
}

export function AreaChart({
  data,
  height = 240,
  color = "var(--primary)",
  showLabels = true,
}: AreaChartProps) {
  const [hoverIndex, setHoverIndex] = useState<number | null>(null);

  const { points, areaPath, linePath, yLabels } = useMemo(() => {
    if (data.length === 0)
      return {
        points: [],
        areaPath: "",
        linePath: "",
        maxClicks: 0,
        yLabels: [],
      };

    const padding = { top: 20, right: 16, bottom: 40, left: 50 };
    const w = 100; // percentage-based width
    const h = height;
    const _innerW = w - padding.left / 6 - padding.right / 6;
    const innerH = h - padding.top - padding.bottom;

    const clicks = data.map((d) => d.clicks);
    const max = Math.max(...clicks, 1);
    const roundedMax = Math.ceil(max / 10) * 10;

    const pts = data.map((d, i) => ({
      x:
        padding.left +
        (i / Math.max(data.length - 1, 1)) *
          (600 - padding.left - padding.right),
      y: padding.top + (1 - d.clicks / roundedMax) * innerH,
      clicks: d.clicks,
      label: new Date(d.bucket).toLocaleDateString("en-US", {
        month: "short",
        day: "numeric",
      }),
      time: new Date(d.bucket).toLocaleTimeString("en-US", {
        hour: "numeric",
      }),
    }));

    const lineSegs = pts
      .map((p, i) => `${i === 0 ? "M" : "L"} ${p.x} ${p.y}`)
      .join(" ");
    const area =
      lineSegs +
      ` L ${pts[pts.length - 1].x} ${padding.top + innerH} L ${pts[0].x} ${padding.top + innerH} Z`;

    const labels = [0, 0.25, 0.5, 0.75, 1].map((frac) => ({
      value: Math.round(roundedMax * (1 - frac)),
      y: padding.top + frac * innerH,
    }));

    return {
      points: pts,
      areaPath: area,
      linePath: lineSegs,
      maxClicks: roundedMax,
      yLabels: labels,
    };
  }, [data, height]);

  if (data.length === 0) {
    return (
      <div
        className="flex items-center justify-center text-muted-foreground text-sm"
        style={{ height }}
      >
        No data available
      </div>
    );
  }

  return (
    <div className="relative w-full" style={{ height }}>
      <svg
        viewBox={`0 0 600 ${height}`}
        className="w-full h-full"
        preserveAspectRatio="none"
        onMouseLeave={() => setHoverIndex(null)}
      >
        <defs>
          <linearGradient id="areaGrad" x1="0" y1="0" x2="0" y2="1">
            <stop offset="0%" stopColor={color} stopOpacity="0.3" />
            <stop offset="100%" stopColor={color} stopOpacity="0.02" />
          </linearGradient>
        </defs>

        {/* Grid lines */}
        {yLabels.map((label) => (
          <line
            key={label.value}
            x1="50"
            y1={label.y}
            x2="584"
            y2={label.y}
            stroke="currentColor"
            strokeOpacity="0.06"
            strokeWidth="1"
          />
        ))}

        {/* Y axis labels */}
        {showLabels &&
          yLabels.map((label) => (
            <text
              key={label.value}
              x="44"
              y={label.y + 4}
              textAnchor="end"
              className="fill-muted-foreground"
              fontSize="11"
              fontFamily="var(--font-mono)"
            >
              {label.value}
            </text>
          ))}

        {/* Area fill */}
        <path d={areaPath} fill="url(#areaGrad)" />

        {/* Line */}
        <path
          d={linePath}
          fill="none"
          stroke={color}
          strokeWidth="2"
          strokeLinecap="round"
          strokeLinejoin="round"
        />

        {/* Hover targets & dots */}
        {points.map((p, i) => (
          // biome-ignore lint/suspicious/noArrayIndexKey: <change during integration>
          <g key={i}>
            {/** biome-ignore lint/a11y/noStaticElementInteractions: <not required> */}
            <rect
              x={i === 0 ? p.x : p.x - (p.x - points[i - 1].x) / 2}
              y="0"
              width={
                i === 0 || i === points.length - 1
                  ? (points[1]?.x ?? 600) - (points[0]?.x ?? 0)
                  : (points[Math.min(i + 1, points.length - 1)].x -
                      points[Math.max(i - 1, 0)].x) /
                    2
              }
              height={height}
              fill="transparent"
              onMouseEnter={() => setHoverIndex(i)}
            />
            {/* Active dot */}
            {hoverIndex === i && (
              <>
                <line
                  x1={p.x}
                  y1="20"
                  x2={p.x}
                  y2={height - 40}
                  stroke={color}
                  strokeOpacity="0.2"
                  strokeWidth="1"
                  strokeDasharray="4 4"
                />
                <circle
                  cx={p.x}
                  cy={p.y}
                  r="5"
                  fill={color}
                  stroke="var(--background)"
                  strokeWidth="2"
                />
              </>
            )}
          </g>
        ))}

        {/* X axis labels */}
        {showLabels &&
          points
            .filter(
              (_, i) =>
                data.length <= 10 ||
                i % Math.ceil(data.length / 7) === 0 ||
                i === data.length - 1,
            )
            .map((p, i) => (
              <text
                // biome-ignore lint/suspicious/noArrayIndexKey: <need to change during integaration>
                key={i}
                x={p.x}
                y={height - 10}
                textAnchor="middle"
                className="fill-muted-foreground"
                fontSize="11"
                fontFamily="var(--font-mono)"
              >
                {data.length <= 24 ? p.time : p.label}
              </text>
            ))}
      </svg>

      {/* Tooltip */}
      {hoverIndex !== null && points[hoverIndex] && (
        <div
          className="absolute pointer-events-none z-10 rounded-lg border bg-popover px-3 py-2 text-sm shadow-lg animate-in fade-in duration-100"
          style={{
            left: `${(points[hoverIndex].x / 600) * 100}%`,
            top: `${points[hoverIndex].y - 10}px`,
            transform: "translate(-50%, -100%)",
          }}
        >
          <p className="font-semibold">{points[hoverIndex].clicks} clicks</p>
          <p className="text-xs text-muted-foreground">
            {data.length <= 24
              ? `${points[hoverIndex].label} ${points[hoverIndex].time}`
              : points[hoverIndex].label}
          </p>
        </div>
      )}
    </div>
  );
}
