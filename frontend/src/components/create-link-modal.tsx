"use client";

import { useState, useCallback } from "react";
import { toast } from "sonner";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Switch } from "@/components/ui/switch";
import type { LinkResponse } from "@/lib/validators/link";
import { createLinkSchema } from "@/lib/validators/link";
import { useCreateLink } from "@/hooks/use-links";
import { linkService } from "@/services/link.service";
import { useDebouncedCallback } from "use-debounce";

export function CreateLinkModal({
  children,
  onCreated,
}: {
  children: React.ReactNode;
  onCreated?: (link: LinkResponse) => void;
}) {
  const [open, setOpen] = useState(false);
  const { mutateAsync: createLink, isPending: loading } = useCreateLink();
  const [result, setResult] = useState<LinkResponse | null>(null);
  const [copied, setCopied] = useState(false);
  const [showUtm, setShowUtm] = useState(false);
  const [errors, setErrors] = useState<Record<string, string>>({});

  const [values, setValues] = useState({
    longUrl: "",
    customAlias: "",
    expiresAt: "",
    redirectType: 302 as 301 | 302,
  });

  const [utm, setUtm] = useState({
    source: "",
    medium: "",
    campaign: "",
  });

  // Alias availability state
  const [aliasStatus, setAliasStatus] = useState<
    "idle" | "checking" | "available" | "taken"
  >("idle");

  const checkAlias = useDebouncedCallback(async (alias: string) => {
    if (!alias || alias.length < 3) {
      setAliasStatus("idle");
      return;
    }

    setAliasStatus("checking");
    try {
      const available = await linkService.checkAliasAvailability(alias);
      setAliasStatus(available ? "available" : "taken");
    } catch (err) {
      setAliasStatus("idle");
    }
  }, 500);

  const buildPreviewUrl = () => {
    if (!values.longUrl) return "";
    try {
      const url = new URL(values.longUrl);
      if (utm.source) url.searchParams.set("utm_source", utm.source);
      if (utm.medium) url.searchParams.set("utm_medium", utm.medium);
      if (utm.campaign) url.searchParams.set("utm_campaign", utm.campaign);
      return url.toString();
    } catch {
      return values.longUrl;
    }
  };

  const handleAliasChange = (alias: string) => {
    setValues((v) => ({ ...v, customAlias: alias }));
    checkAlias(alias);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setErrors({});

    const finalUrl = showUtm ? buildPreviewUrl() : values.longUrl;
    const payload = { ...values, longUrl: finalUrl };

    const parsed = createLinkSchema.safeParse(payload);
    if (!parsed.success) {
      const fieldErrors: Record<string, string> = {};
      for (const issue of parsed.error.issues) {
        const path = issue.path.join(".");
        if (!fieldErrors[path]) fieldErrors[path] = issue.message;
      }
      setErrors(fieldErrors);
      return;
    }

    if (aliasStatus === "taken") {
      setErrors({ customAlias: "This alias is already taken" });
      return;
    }

    try {
      const created = await createLink({
        longUrl: finalUrl,
        customAlias: values.customAlias || undefined,
        redirectType: values.redirectType,
        expiresAt: values.expiresAt || undefined,
      });
      setResult(created);
      onCreated?.(created);
    } catch (err) {
      // Error handled by hook
    }
  };

  const handleCopy = async () => {
    if (!result) return;
    await navigator.clipboard.writeText(result.shortUrl);
    setCopied(true);
    toast.success("Copied to clipboard!");
    setTimeout(() => setCopied(false), 2000);
  };

  const handleClose = (isOpen: boolean) => {
    setOpen(isOpen);
    if (!isOpen) {
      // Reset state when closing
      setTimeout(() => {
        setResult(null);
        setCopied(false);
        setErrors({});
        setAliasStatus("idle");
        setShowUtm(false);
        setValues({
          longUrl: "",
          customAlias: "",
          expiresAt: "",
          redirectType: 302,
        });
        setUtm({ source: "", medium: "", campaign: "" });
      }, 200);
    }
  };

  return (
    <Dialog open={open} onOpenChange={handleClose}>
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent className="sm:max-w-lg max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle className="text-xl font-bold">
            {result ? "Link Created! 🎉" : "Create New Link"}
          </DialogTitle>
        </DialogHeader>

        {/* Success State */}
        {result ? (
          <div className="space-y-4 animate-in fade-in slide-in-from-bottom-2 duration-300">
            <div className="rounded-xl bg-muted/50 p-4 space-y-3">
              <div>
                <p className="text-xs text-muted-foreground mb-1">
                  Original URL
                </p>
                <p className="text-sm truncate">{result.longUrl}</p>
              </div>
              <div>
                <p className="text-xs text-muted-foreground mb-1">Short URL</p>
                <p className="text-lg font-bold text-primary font-mono">
                  {result.shortUrl}
                </p>
              </div>
            </div>

            <Button
              onClick={handleCopy}
              className="w-full h-12 text-base font-semibold rounded-xl"
              variant={copied ? "secondary" : "default"}
            >
              {copied ? (
                <span className="flex items-center gap-2">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="18"
                    height="18"
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
                    width="18"
                    height="18"
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
                  Copy to Clipboard
                </span>
              )}
            </Button>

            <Button
              variant="outline"
              className="w-full rounded-xl"
              onClick={() => handleClose(false)}
            >
              Done
            </Button>
          </div>
        ) : (
          /* Create Form */
          <form onSubmit={handleSubmit} className="space-y-5">
            {/* Long URL */}
            <div className="space-y-2">
              <Label htmlFor="create-url">Destination URL</Label>
              <Input
                id="create-url"
                type="url"
                placeholder="https://example.com/your-long-url"
                value={values.longUrl}
                onChange={(e) =>
                  setValues((v) => ({ ...v, longUrl: e.target.value }))
                }
                autoFocus
              />
              {errors.longUrl && (
                <p className="text-xs text-destructive">{errors.longUrl}</p>
              )}
            </div>

            {/* Custom Alias */}
            <div className="space-y-2">
              <Label htmlFor="create-alias">
                Custom Alias{" "}
                <span className="text-muted-foreground font-normal">
                  (optional)
                </span>
              </Label>
              <div className="relative">
                <Input
                  id="create-alias"
                  placeholder="my-custom-link"
                  value={values.customAlias}
                  onChange={(e) => handleAliasChange(e.target.value)}
                  className="pr-10"
                />
                {aliasStatus !== "idle" && (
                  <div className="absolute right-3 top-1/2 -translate-y-1/2">
                    {aliasStatus === "checking" && (
                      <svg
                        className="h-4 w-4 animate-spin text-muted-foreground"
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
                          d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"
                        />
                      </svg>
                    )}
                    {aliasStatus === "available" && (
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
                        className="text-green-500"
                      >
                        <polyline points="20 6 9 17 4 12" />
                      </svg>
                    )}
                    {aliasStatus === "taken" && (
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
                        className="text-destructive"
                      >
                        <path d="M18 6 6 18" />
                        <path d="m6 6 12 12" />
                      </svg>
                    )}
                  </div>
                )}
              </div>
              {errors.customAlias && (
                <p className="text-xs text-destructive">{errors.customAlias}</p>
              )}
            </div>

            {/* Redirect Type */}
            <div className="flex items-center justify-between rounded-lg bg-muted/50 p-3">
              <div>
                <p className="text-sm font-medium">Permanent Redirect (301)</p>
                <p className="text-xs text-muted-foreground">
                  Cached by browsers. Use for permanent destinations.
                </p>
              </div>
              <Switch
                checked={values.redirectType === 301}
                onCheckedChange={(checked) =>
                  setValues((v) => ({
                    ...v,
                    redirectType: checked ? 301 : 302,
                  }))
                }
              />
            </div>

            {/* Expiry */}
            <div className="space-y-2">
              <Label htmlFor="create-expires">
                Expiration Date{" "}
                <span className="text-muted-foreground font-normal">
                  (optional)
                </span>
              </Label>
              <Input
                id="create-expires"
                type="datetime-local"
                value={values.expiresAt}
                onChange={(e) =>
                  setValues((v) => ({ ...v, expiresAt: e.target.value }))
                }
              />
            </div>

            {/* UTM Builder Toggle */}
            <div>
              <button
                type="button"
                onClick={() => setShowUtm(!showUtm)}
                className="flex items-center gap-2 text-sm font-medium text-primary hover:underline underline-offset-4"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="14"
                  height="14"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  strokeWidth="2"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  className={`transition-transform ${showUtm ? "rotate-90" : ""}`}
                >
                  <path d="m9 18 6-6-6-6" />
                </svg>
                UTM Parameters
              </button>

              {showUtm && (
                <div className="mt-3 space-y-3 animate-in fade-in slide-in-from-top-2 duration-200">
                  <Input
                    placeholder="Source (e.g. twitter)"
                    value={utm.source}
                    onChange={(e) =>
                      setUtm((u) => ({ ...u, source: e.target.value }))
                    }
                  />
                  <Input
                    placeholder="Medium (e.g. social)"
                    value={utm.medium}
                    onChange={(e) =>
                      setUtm((u) => ({ ...u, medium: e.target.value }))
                    }
                  />
                  <Input
                    placeholder="Campaign (e.g. launch-2026)"
                    value={utm.campaign}
                    onChange={(e) =>
                      setUtm((u) => ({ ...u, campaign: e.target.value }))
                    }
                  />
                  {(utm.source || utm.medium || utm.campaign) && (
                    <div className="rounded-lg bg-muted/50 p-3">
                      <p className="text-xs text-muted-foreground mb-1">
                        Preview URL
                      </p>
                      <p className="text-xs font-mono break-all">
                        {buildPreviewUrl()}
                      </p>
                    </div>
                  )}
                </div>
              )}
            </div>

            <Button
              type="submit"
              className="w-full h-11 font-semibold rounded-xl"
              disabled={loading || !values.longUrl.trim()}
            >
              {loading ? "Creating..." : "Create Short Link"}
            </Button>
          </form>
        )}
      </DialogContent>
    </Dialog>
  );
}
