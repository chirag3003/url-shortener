import { Skeleton } from "@/components/ui/skeleton";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

export function LinksTableSkeleton() {
  return (
    <div className="rounded-xl border bg-card overflow-hidden">
      <Table>
        <TableHeader className="bg-muted/30">
          <TableRow>
            <TableHead className="w-[200px]">Short Link</TableHead>
            <TableHead className="max-w-[400px]">Original URL</TableHead>
            <TableHead className="w-[100px] text-center">Status</TableHead>
            <TableHead className="w-[100px] text-right">Clicks</TableHead>
            <TableHead className="w-[100px] text-right">Actions</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {Array.from({ length: 5 }).map((_, i) => (
            <TableRow key={i}>
              <TableCell>
                <Skeleton className="h-5 w-32" />
              </TableCell>
              <TableCell>
                <Skeleton className="h-4 w-full max-w-[300px]" />
              </TableCell>
              <TableCell>
                <div className="flex justify-center">
                  <Skeleton className="h-6 w-16 rounded-full" />
                </div>
              </TableCell>
              <TableCell>
                <div className="flex justify-end">
                  <Skeleton className="h-5 w-12" />
                </div>
              </TableCell>
              <TableCell>
                <div className="flex justify-end gap-2">
                  <Skeleton className="h-8 w-8 rounded-md" />
                  <Skeleton className="h-8 w-8 rounded-md" />
                </div>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
}
