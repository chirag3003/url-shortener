"use client";

import { useRef, useState } from "react";
import { toast } from "sonner";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import type { LinkResponse } from "@/lib/validators/link";
import { createLinkSchema } from "@/lib/validators/link";
import { useAuth } from "@/hooks/use-auth";
import { AuthModal } from "./auth-modal";
import { useQuickShorten } from "@/hooks/use-links";

export function QuickShortener() {
  const { isAuthenticated } = useAuth();
  const { mutateAsync: quickShorten, isPending: loading } = useQuickShorten();
  const [url, setUrl] = useState("");
  const [result, setResult] = useState<LinkResponse | null>(null);
  const [error, setError] = useState("");
  const [copied, setCopied] = useState(false);
  const inputRef = useRef<HTMLInputElement>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!isAuthenticated) return;

    setError("");
    setResult(null);

    const parsed = createLinkSchema.safeParse({ longUrl: url });
    if (!parsed.success) {
      setError(parsed.error.issues[0]?.message ?? "Invalid URL");
      return;
    }

    try {
      const data = await quickShorten(url);
      setResult(data);
      toast.success("Link shortened successfully!");
    } catch (err) {
      // Error handled by hook toast
    }
  };

  const handleCopy = async () => {
    if (!result) return;
    await navigator.clipboard.writeText(result.shortUrl);
    setCopied(true);
    toast.success("Copied to clipboard!");
    setTimeout(() => setCopied(false), 2000);
  };

  const handleReset = () => {
    setResult(null);
    setUrl("");
    setError("");
    setCopied(false);
    setTimeout(() => inputRef.current?.focus(), 100);
  };

  return (
    <div className="w-full max-w-2xl mx-auto">
      {/* Input form */}
      {!result && (
        <AuthModal>
          <form
            onSubmit={(e) => {
              if (isAuthenticated) {
                handleSubmit(e);
              }
            }}
            className="relative"
          >
            <div className="relative group">
              <div className="absolute -inset-0.5 rounded-2xl bg-gradient-to-r from-primary/60 via-accent/60 to-secondary/60 opacity-0 blur-lg transition-opacity duration-500 group-focus-within:opacity-100" />
              <div className="relative flex items-center gap-2 rounded-2xl border bg-card p-2 shadow-lg transition-shadow group-focus-within:shadow-xl">
                <div className="flex items-center pl-3 text-muted-foreground">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="20"
                    height="20"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    strokeWidth="2"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                  >
                    <path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71" />
                    <path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71" />
                  </svg>
                </div>
                <Input
                  ref={inputRef}
                  type="url"
                  placeholder="Paste your long URL here..."
                  value={url}
                  onChange={(e) => {
                    setUrl(e.target.value);
                    if (error) setError("");
                  }}
                  onFocus={(e) => {
                    if (!isAuthenticated) {
                      e.target.blur();
                    }
                  }}
                  className="flex-1 border-0 bg-transparent text-base ring-0 shadow-none focus-visible:ring-0 focus-visible:ring-offset-0 placeholder:text-muted-foreground/60"
                  autoFocus
                />
                <Button
                  type={isAuthenticated ? "submit" : "button"}
                  size="lg"
                  disabled={loading || (isAuthenticated && !url.trim())}
                  className="rounded-xl px-6 font-semibold transition-all hover:scale-[1.02] active:scale-[0.98]"
                >
                  {loading ? (
                    <span className="flex items-center gap-2">
                      <svg
                        className="h-4 w-4 animate-spin"
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                      >
                        <circle
                          className="opacity-25"
                          cx="12"
                          cy="12"
                          r="10"
                          stroke="currentColor"
                          strokeWidth="4"
                        />
                        <path
                          className="opacity-75"
                          fill="currentColor"
                          d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                        />
                      </svg>
                      Shortening...
                    </span>
                  ) : (
                    "Shorten"
                  )}
                </Button>
              </div>
            </div>
            {error && (
              <p className="mt-3 text-sm text-destructive text-center animate-in fade-in slide-in-from-top-1 duration-200">
                {error}
              </p>
            )}
          </form>
        </AuthModal>
      )}

      {/* Result card */}
      {result && (
        <div className="animate-in fade-in slide-in-from-bottom-4 duration-500">
          <div className="rounded-2xl border bg-card p-6 shadow-lg">
            <div className="flex items-start justify-between gap-4 mb-4">
              <div className="min-w-0 flex-1">
                <p className="text-sm text-muted-foreground truncate">
                  {result.longUrl}
                </p>
              </div>
              <span className="inline-flex items-center rounded-full bg-green-500/10 px-2.5 py-0.5 text-xs font-medium text-green-600 dark:text-green-400 whitespace-nowrap">
                ✓ Created
              </span>
            </div>

            <div className="flex items-center gap-3 rounded-xl bg-muted/50 p-4">
              <div className="flex-1 min-w-0">
                <p className="text-lg font-semibold text-primary truncate font-mono">
                  {result.shortUrl}
                </p>
              </div>
              <Button
                onClick={handleCopy}
                variant={copied ? "secondary" : "default"}
                size="lg"
                className="rounded-xl px-6 font-semibold shrink-0 transition-all hover:scale-[1.02] active:scale-[0.98]"
              >
                {copied ? (
                  <span className="flex items-center gap-2">
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      width="16"
                      height="16"
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      strokeWidth="2.5"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                    >
                      <polyline points="20 6 9 17 4 12" />
                    </svg>
                    Copied!
                  </span>
                ) : (
                  <span className="flex items-center gap-2">
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
                      <rect width="14" height="14" x="8" y="8" rx="2" ry="2" />
                      <path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2" />
                    </svg>
                    Copy
                  </span>
                )}
              </Button>
            </div>

            <div className="mt-4 flex justify-center">
              {/** biome-ignore lint/a11y/useButtonType: <later> */}
              <button
                onClick={handleReset}
                className="text-sm text-muted-foreground hover:text-foreground transition-colors underline underline-offset-4"
              >
                Shorten another link
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
