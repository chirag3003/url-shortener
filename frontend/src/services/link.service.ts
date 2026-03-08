import apiClient from "@/lib/api/client";
import type {
  AliasAvailabilityResponse,
  CreateLinkInput,
  LinkResponse,
  ListLinksQuery,
  PaginatedLinksResponse,
  UpdateLinkInput,
} from "@/lib/validators/link";

export const linkService = {
  async createLink(data: CreateLinkInput): Promise<LinkResponse> {
    const response = await apiClient.post<LinkResponse>("/api/v1/links", data);
    return response.data;
  },

  async quickShorten(longUrl: string): Promise<LinkResponse> {
    const response = await apiClient.post<LinkResponse>("/api/v1/links/quick", {
      longUrl,
    });
    return response.data;
  },

  async listLinks(query?: ListLinksQuery): Promise<PaginatedLinksResponse> {
    const response = await apiClient.get<PaginatedLinksResponse>("/api/v1/links", {
      params: query,
    });
    return response.data;
  },

  async getLink(id: string): Promise<LinkResponse> {
    const response = await apiClient.get<LinkResponse>(`/api/v1/links/${id}`);
    return response.data;
  },

  async updateLink(id: string, data: UpdateLinkInput): Promise<LinkResponse> {
    const response = await apiClient.patch<LinkResponse>(
      `/api/v1/links/${id}`,
      data
    );
    return response.data;
  },

  async deleteLink(id: string): Promise<void> {
    await apiClient.delete(`/api/v1/links/${id}`);
  },

  async checkAliasAvailability(
    alias: string
  ): Promise<AliasAvailabilityResponse> {
    const response = await apiClient.get<AliasAvailabilityResponse>(
      "/api/v1/links/alias-availability",
      {
        params: { alias },
      }
    );
    return response.data;
  },
};
