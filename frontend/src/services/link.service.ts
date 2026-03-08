import apiClient from "@/lib/api/client";
import type {
  AliasAvailabilityResponse,
  CreateLinkInput,
  LinkResponse,
  ListLinksQuery,
  PaginatedLinksResponse,
  UpdateLinkInput,
} from "@/lib/validators/link";

export interface ApiLinkResponse {
  data: LinkResponse;
}

export interface ApiLinkListResponse {
  data: PaginatedLinksResponse;
}

export interface ApiAliasAvailabilityResponse {
  data: AliasAvailabilityResponse;
}


export const linkService = {
  async createLink(data: CreateLinkInput): Promise<LinkResponse> {
    const response = await apiClient.post<ApiLinkResponse>("/api/v1/links", data);
    return response.data.data;
  },

  async quickShorten(longUrl: string): Promise<LinkResponse> {
    const response = await apiClient.post<ApiLinkResponse>("/api/v1/links/quick", {
      longUrl,
    });
    return response.data.data;
  },

  async listLinks(query?: ListLinksQuery): Promise<PaginatedLinksResponse> {
    const response = await apiClient.get<ApiLinkListResponse>("/api/v1/links", {
      params: query,
    });
    return response.data.data;
  },

  async getLink(id: string): Promise<LinkResponse> {
    const response = await apiClient.get<ApiLinkResponse>(`/api/v1/links/${id}`);
    return response.data.data;
  },

  async updateLink(id: string, data: UpdateLinkInput): Promise<LinkResponse> {
    const response = await apiClient.patch<ApiLinkResponse>(
      `/api/v1/links/${id}`,
      data
    );
    return response.data.data;
  },

  async deleteLink(id: string): Promise<void> {
    await apiClient.delete(`/api/v1/links/${id}`);
  },

  async checkAliasAvailability(
    alias: string
  ): Promise<AliasAvailabilityResponse> {
    const response = await apiClient.get<ApiAliasAvailabilityResponse>(
      "/api/v1/links/alias-availability",
      {
        params: { alias },
      }
    );
    return response.data.data;
  },
};
