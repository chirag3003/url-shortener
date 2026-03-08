"use client";

import { useState, useEffect } from "react";
import { toast } from "sonner";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useAuth } from "@/hooks/use-auth";
import { authService } from "@/services/auth.service";

export function ProfileView() {
  const { user, refreshUser } = useAuth();
  const [name, setName] = useState(user?.name ?? "");
  const [email, setEmail] = useState(user?.email ?? "");
  const [saving, setSaving] = useState(false);

  useEffect(() => {
    if (user) {
      setName(user.name);
      setEmail(user.email);
    }
  }, [user]);

  const handleSave = async () => {
    if (!name.trim() || !email.trim()) {
      toast.error("Name and email are required");
      return;
    }

    setSaving(true);
    try {
      await authService.updateMe({ name, email });
      await refreshUser();
      toast.success("Profile updated!");
    } catch (error: any) {
      toast.error(error.response?.data?.error || "Failed to update profile");
    } finally {
      setSaving(false);
    }
  };

  return (
    <div className="space-y-6 animate-in fade-in duration-300">
      <div>
        <h2 className="text-xl font-semibold tracking-tight">Public Profile</h2>
        <p className="text-sm text-muted-foreground mt-1">
          This is how others will see you on the site.
        </p>
      </div>
      <div className="h-px bg-border/60" />

      <div className="space-y-8 max-w-xl">
        <div className="space-y-3">
          <Label htmlFor="profile-name" className="text-sm font-medium">
            Name
          </Label>
          <Input
            id="profile-name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="max-w-md h-10"
          />
          <p className="text-[13px] text-muted-foreground">
            Your name may appear around Snip.ly where you are mentioned.
          </p>
        </div>

        <div className="space-y-3">
          <Label htmlFor="profile-email" className="text-sm font-medium">
            Email Address
          </Label>
          <Input
            id="profile-email"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="max-w-md h-10"
          />
          <p className="text-[13px] text-muted-foreground">
            We will use this email address to contact you about your account.
          </p>
        </div>

        <Button
          onClick={handleSave}
          disabled={saving}
          className="rounded-lg font-medium"
        >
          {saving ? "Saving..." : "Update Profile"}
        </Button>
      </div>
    </div>
  );
}
