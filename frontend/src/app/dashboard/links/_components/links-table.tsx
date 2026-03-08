"use client";

import Link from "next/link";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { mockClickCounts } from "@/lib/mock-data";
import type { LinkResponse } from "@/lib/validators/link";

function formatDate(iso: string) {
  return new Date(iso).toLocaleDateString("en-US", {
    month: "short",
    day: "numeric",
    year: "numeric",
  });
}

function truncateUrl(url: string, maxLen = 50) {
  if (url.length <= maxLen) return url;
  return `${url.substring(0, maxLen)}…`;
}

interface LinksTableProps {
  paginated: LinkResponse[];
  handleCopy: (shortUrl: string) => void;
  setDeleteTarget: (link: LinkResponse) => void;
}

export function LinksTable({
  paginated,
  handleCopy,
  setDeleteTarget,
}: LinksTableProps) {
  return (
    <div className="rounded-xl border overflow-hidden bg-card/50">
      <Table>
        <TableHeader>
          <TableRow className="bg-muted/30">
            <TableHead className="font-semibold">Short Link</TableHead>
            <TableHead className="font-semibold hidden sm:table-cell">
              Destination
            </TableHead>
            <TableHead className="font-semibold">Clicks</TableHead>
            <TableHead className="font-semibold hidden md:table-cell">
              Created
            </TableHead>
            <TableHead className="font-semibold">Status</TableHead>
            <TableHead className="w-[50px]" />
          </TableRow>
        </TableHeader>
        <TableBody>
          {paginated.map((link) => (
            <TableRow key={link.id} className="group">
              <TableCell>
                <p className="font-semibold text-primary font-mono text-sm">
                  /{link.shortCode}
                </p>
              </TableCell>
              <TableCell className="hidden sm:table-cell max-w-[250px]">
                <p className="text-sm text-muted-foreground truncate">
                  {truncateUrl(link.longUrl)}
                </p>
              </TableCell>
              <TableCell>
                <p className="font-semibold text-sm">
                  {(mockClickCounts[link.id] ?? 0).toLocaleString()}
                </p>
              </TableCell>
              <TableCell className="hidden md:table-cell">
                <p className="text-sm text-muted-foreground">
                  {formatDate(link.createdAt)}
                </p>
              </TableCell>
              <TableCell>
                <Badge
                  variant={link.isActive ? "default" : "secondary"}
                  className="text-[10px]"
                >
                  {link.isActive ? "Active" : "Inactive"}
                </Badge>
              </TableCell>
              <TableCell>
                <DropdownMenu>
                  <DropdownMenuTrigger asChild>
                    <Button
                      variant="ghost"
                      size="icon"
                      className="h-8 w-8 opacity-0 group-hover:opacity-100 transition-opacity"
                    >
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="16"
                        height="16"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        strokeWidth="2"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                      >
                        <circle cx="12" cy="12" r="1" />
                        <circle cx="12" cy="5" r="1" />
                        <circle cx="12" cy="19" r="1" />
                      </svg>
                    </Button>
                  </DropdownMenuTrigger>
                  <DropdownMenuContent align="end" className="w-44">
                    <DropdownMenuItem asChild>
                      <Link href={`/dashboard/links/${link.id}/analytics`}>
                        View Analytics
                      </Link>
                    </DropdownMenuItem>
                    <DropdownMenuItem onClick={() => handleCopy(link.shortUrl)}>
                      Copy Link
                    </DropdownMenuItem>
                    <DropdownMenuSeparator />
                    <DropdownMenuItem
                      className="text-destructive focus:text-destructive"
                      onClick={() => setDeleteTarget(link)}
                    >
                      Delete
                    </DropdownMenuItem>
                  </DropdownMenuContent>
                </DropdownMenu>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
}
