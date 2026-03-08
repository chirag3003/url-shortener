"use client";

import { useState } from "react";
import { AreaChart } from "@/components/charts/area-chart";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Tabs, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Skeleton } from "@/components/ui/skeleton";

interface ClicksChartProps {
  timeSeriesData: {
    "24h": any[];
    "7d": any[];
    "30d": any[];
  };
}

export function ClicksChart({ timeSeriesData }: ClicksChartProps) {
  const [window, setWindow] = useState<"24h" | "7d" | "30d">("30d");

  const currentData = timeSeriesData[window];
  const isLoading = !currentData || currentData.length === 0;

  return (
    <Card>
      <CardHeader className="flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4">
        <CardTitle>Click Traffic</CardTitle>
        <Tabs
          value={window}
          onValueChange={(v) => setWindow(v as "24h" | "7d" | "30d")}
        >
          <TabsList>
            <TabsTrigger value="24h" className="text-xs">
              24 Hours
            </TabsTrigger>
            <TabsTrigger value="7d" className="text-xs">
              7 Days
            </TabsTrigger>
            <TabsTrigger value="30d" className="text-xs">
              30 Days
            </TabsTrigger>
          </TabsList>
        </Tabs>
      </CardHeader>
      <CardContent>
        {isLoading ? (
          <div className="flex items-end gap-2 h-[280px] w-full pt-8">
            {[1, 2, 3, 4, 5, 6, 7, 8, 9, 10].map((i) => (
              <Skeleton 
                key={i} 
                className="flex-1 rounded-t-sm" 
                style={{ height: `${Math.random() * 60 + 20}%` }}
              />
            ))}
          </div>
        ) : (
          <AreaChart data={currentData} height={280} />
        )}
      </CardContent>
    </Card>
  );
}
