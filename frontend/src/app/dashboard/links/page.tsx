"use client";

import { useMemo, useState } from "react";
import { toast } from "sonner";
import { CreateLinkModal } from "@/components/create-link-modal";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Skeleton } from "@/components/ui/skeleton";
import { mockLinks } from "@/lib/mock-data";
import type { LinkResponse } from "@/lib/validators/link";
import { LinksHeader } from "./_components/links-header";
import { LinksTable } from "./_components/links-table";

// ── Empty State ──────────────────────────────────────────────────────────

function EmptyState() {
  return (
    <div className="flex flex-col items-center justify-center py-16 px-4 text-center">
      <div className="mb-6 flex h-20 w-20 items-center justify-center rounded-2xl bg-muted">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="40"
          height="40"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          strokeWidth="1"
          strokeLinecap="round"
          strokeLinejoin="round"
          className="text-muted-foreground"
        >
          <path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71" />
          <path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71" />
        </svg>
      </div>
      <h3 className="text-lg font-semibold">No links yet</h3>
      <p className="mt-1 text-sm text-muted-foreground max-w-sm">
        You haven&apos;t shortened any links yet. Create your first one and
        start tracking clicks!
      </p>
      <CreateLinkModal onCreated={() => {}}>
        <Button className="mt-6 rounded-xl font-semibold">
          Create Your First Link
        </Button>
      </CreateLinkModal>
    </div>
  );
}

// ── Skeleton Loader ──────────────────────────────────────────────────────

function TableSkeleton() {
  return (
    <div className="space-y-3">
      {Array.from({ length: 5 }).map((_, i) => (
        <div key={i} className="flex items-center gap-4 py-3">
          <Skeleton className="h-4 w-[200px]" />
          <Skeleton className="h-4 w-[120px]" />
          <Skeleton className="h-4 w-[80px]" />
          <Skeleton className="h-4 w-[60px]" />
          <Skeleton className="h-4 w-[28px]" />
        </div>
      ))}
    </div>
  );
}

// ── Page ─────────────────────────────────────────────────────────────────

export default function LinksPage() {
  const [search, setSearch] = useState("");
  const [links, setLinks] = useState<LinkResponse[]>(mockLinks);
  const [deleteTarget, setDeleteTarget] = useState<LinkResponse | null>(null);
  const [page, setPage] = useState(1);
  const perPage = 10;

  // Filter links by search
  const filtered = useMemo(() => {
    if (!search.trim()) return links;
    const q = search.toLowerCase();
    return links.filter(
      (link) =>
        link.shortCode.toLowerCase().includes(q) ||
        link.longUrl.toLowerCase().includes(q),
    );
  }, [links, search]);

  // Paginate
  const totalPages = Math.ceil(filtered.length / perPage);
  const paginated = filtered.slice((page - 1) * perPage, page * perPage);

  const handleCopy = (shortUrl: string) => {
    navigator.clipboard.writeText(shortUrl);
    toast.success("Link copied!");
  };

  const handleDelete = (link: LinkResponse) => {
    setLinks((prev) => prev.filter((l) => l.id !== link.id));
    setDeleteTarget(null);
    toast.success(`Link /${link.shortCode} deleted.`);
  };

  const handleCreated = (newLink: LinkResponse) => {
    setLinks((prev) => [newLink, ...prev]);
  };

  return (
    <div className="space-y-6">
      {/* Header */}
      <LinksHeader onCreated={handleCreated} />

      {/* Search */}
      <div className="relative max-w-md">
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
          className="absolute left-3 top-1/2 -translate-y-1/2 text-muted-foreground"
        >
          <circle cx="11" cy="11" r="8" />
          <path d="m21 21-4.3-4.3" />
        </svg>
        <Input
          placeholder="Search by URL or alias..."
          value={search}
          onChange={(e) => {
            setSearch(e.target.value);
            setPage(1);
          }}
          className="pl-10 h-10"
        />
      </div>

      {/* Table */}
      {links.length === 0 ? (
        <EmptyState />
      ) : filtered.length === 0 ? (
        <div className="text-center py-12 text-muted-foreground">
          No links match your search.
        </div>
      ) : (
        <>
          <LinksTable
            paginated={paginated}
            handleCopy={handleCopy}
            setDeleteTarget={setDeleteTarget}
          />

          {/* Pagination */}
          {totalPages > 1 && (
            <div className="flex items-center justify-between">
              <p className="text-sm text-muted-foreground">
                Showing {(page - 1) * perPage + 1}–
                {Math.min(page * perPage, filtered.length)} of {filtered.length}{" "}
                links
              </p>
              <div className="flex gap-2">
                <Button
                  variant="outline"
                  size="sm"
                  disabled={page === 1}
                  onClick={() => setPage((p) => p - 1)}
                >
                  Previous
                </Button>
                <Button
                  variant="outline"
                  size="sm"
                  disabled={page === totalPages}
                  onClick={() => setPage((p) => p + 1)}
                >
                  Next
                </Button>
              </div>
            </div>
          )}
        </>
      )}

      {/* Delete Confirmation Dialog */}
      <Dialog
        open={!!deleteTarget}
        onOpenChange={(open) => !open && setDeleteTarget(null)}
      >
        <DialogContent className="sm:max-w-sm">
          <DialogHeader>
            <DialogTitle>Delete Link</DialogTitle>
            <DialogDescription>
              Are you sure you want to delete{" "}
              <span className="font-semibold text-foreground font-mono">
                /{deleteTarget?.shortCode}
              </span>
              ? This action cannot be undone, and all analytics data will be
              lost.
            </DialogDescription>
          </DialogHeader>
          <div className="flex gap-3 mt-4">
            <Button
              variant="outline"
              className="flex-1"
              onClick={() => setDeleteTarget(null)}
            >
              Cancel
            </Button>
            <Button
              variant="destructive"
              className="flex-1"
              onClick={() => deleteTarget && handleDelete(deleteTarget)}
            >
              Delete
            </Button>
          </div>
        </DialogContent>
      </Dialog>
    </div>
  );
}
