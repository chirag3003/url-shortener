"use client";

import { CreateLinkModal } from "@/components/create-link-modal";
import { Button } from "@/components/ui/button";
import type { LinkResponse } from "@/lib/validators/link";

interface LinksHeaderProps {
  onCreated: (link: LinkResponse) => void;
}

export function LinksHeader({ onCreated }: LinksHeaderProps) {
  return (
    <div className="flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4">
      <div>
        <h1 className="text-2xl font-bold tracking-tight sm:text-3xl">Links</h1>
        <p className="mt-1 text-muted-foreground">
          Manage all your shortened links.
        </p>
      </div>
      <CreateLinkModal onCreated={onCreated}>
        <Button className="rounded-xl font-semibold shrink-0">
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
            className="mr-2"
          >
            <path d="M5 12h14" />
            <path d="M12 5v14" />
          </svg>
          Create Link
        </Button>
      </CreateLinkModal>
    </div>
  );
}
