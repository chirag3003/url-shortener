import { linkService } from "@/services/link.service";
import type {
  CreateLinkInput,
  ListLinksQuery,
  UpdateLinkInput,
} from "@/lib/validators/link";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";

export function useLinks(query?: ListLinksQuery) {
  return useQuery({
    queryKey: ["links", query],
    queryFn: () => linkService.listLinks(query),
    staleTime: 1000 * 60, // 1 minute
  });
}

export function useLink(id: string) {
  return useQuery({
    queryKey: ["link", id],
    queryFn: () => linkService.getLink(id),
    enabled: !!id,
  });
}

export function useCreateLink() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: CreateLinkInput) => linkService.createLink(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["links"] });
      toast.success("Link created successfully");
    },
    onError: (error: any) => {
      toast.error(error.response?.data?.error || "Failed to create link");
    },
  });
}

export function useQuickShorten() {
  return useMutation({
    mutationFn: (longUrl: string) => linkService.quickShorten(longUrl),
    onSuccess: () => {
      toast.success("Link shortened successfully");
    },
    onError: (error: any) => {
      toast.error(error.response?.data?.error || "Failed to shorten link");
    },
  });
}

export function useUpdateLink() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateLinkInput }) =>
      linkService.updateLink(id, data),
    onSuccess: (data, variables) => {
      queryClient.invalidateQueries({ queryKey: ["links"] });
      queryClient.invalidateQueries({ queryKey: ["link", variables.id] });
    },
    onError: (error: any) => {
      toast.error(error.response?.data?.error || "Failed to update link");
    },
  });
}

export function useDeleteLink() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => linkService.deleteLink(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["links"] });
      toast.success("Link deleted successfully");
    },
    onError: (error: any) => {
      toast.error(error.response?.data?.error || "Failed to delete link");
    },
  });
}
