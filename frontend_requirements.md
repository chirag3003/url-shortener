# UI/UX Feature Specification: URL Shortener

This document outlines the user interface components, user experience flows, and interaction states required for the full-stack URL shortener platform.

## 1. Landing & Authentication Flows
**Goal:** Provide immediate value to unauthenticated users while driving conversions to registered accounts.

* **The Public "Quick Shortener" Widget:**
  * **UI:** A massive, centrally located input field on the landing page.
  * **UX:** Pressing `Enter` instantly generates a short link. 
  * **State Management:** Shows a pulsing skeleton loader during the network request.
  * **Conversion Hook:** Once generated, displays a toast notification: "Link created! Sign up to track clicks."
* **Authentication Modals (Login/Signup):**
  * **UI:** Clean, centralized card layout.
  * **UX:** Support for email/password and OAuth (Google/GitHub). Includes inline validation for password strength and email formatting before submission.

## 2. The Command Center (Main Dashboard)
**Goal:** A friction-free workspace for link generation and management.



* **Global Action Bar:**
  * **UI:** Sticky header with a persistent "Create New Link" button.
  * **UX:** Accessible from anywhere in the app to prevent users from needing to navigate back to the home page to shorten a link.
* **Advanced Link Creation Modal:**
  * **Core Input:** Auto-focuses on the Long URL field upon opening.
  * **Custom Alias Field:** Real-time debounced validation (checks if `domain.com/custom` is already taken while the user types, showing a red/green status icon).
  * **UTM Builder (Collapsible):** Simple text inputs for Source, Medium, and Campaign that automatically append to the preview URL.
  * **Expiration Picker:** A standard calendar UI component to select a self-destruct date.
* **Success State & Handoff:**
  * **UI:** A temporary overlay or expanded card once the link is created.
  * **UX:** Features a massive "Copy to Clipboard" button (with a tooltip that changes from "Copy" to "Copied!" on click) and a "Download QR Code" button.

## 3. The Links Data Grid
**Goal:** Efficiently organize and filter hundreds of generated links.

* **The Table View:**
  * **UI:** A paginated, responsive data table displaying: Original URL (truncated), Short URL, Creation Date, and Total Clicks.
  * **UX States:** * *Empty State:* A friendly illustration with a clear call-to-action ("You haven't shortened any links yet. Create your first one!").
    * *Loading State:* Uses skeleton rows instead of a generic spinner to reduce perceived latency.
* **Filtering & Search:**
  * **UX:** A dynamic search bar that filters the table in real-time by keyword, custom alias, or date range.
* **Row-Level Actions:**
  * **UI:** A "three-dot" context menu on each row.
  * **Actions:** "View Analytics", "Copy Link", "Edit Destination" (if supported), and "Archive/Delete" (with a confirmation modal to prevent accidental deletion).

## 4. Analytics & Telemetry Dashboard
**Goal:** Provide clear, actionable data visualization for a specific short link.



* **Macro Metrics (Top Cards):**
  * **UI:** Large, high-contrast typography displaying Total Clicks and Unique Visitors. Includes a trend indicator (e.g., a green arrow showing "+15% this week").
* **Time-Series Traffic Chart:**
  * **UI:** An interactive line or area chart.
  * **UX:** Users can hover over data points to see the exact number of clicks for that specific day/hour. Includes a toggle group to switch the X-axis between 24 Hours, 7 Days, and 30 Days.
* **Categorical Breakdown Panels:**
  * **Geospatial Map:** A choropleth map where countries are shaded based on traffic density. Hovering reveals country-specific click counts.
  * **Referrer List:** A clean, ranked list showing traffic sources (e.g., Twitter, Google, Direct) alongside their respective favicon.
  * **Device & Browser Donut Charts:** Simple, color-coded rings breaking down the hardware and software used by visitors.

## 5. User Settings & API Management
**Goal:** Allow power users to manage their integration credentials securely.

* **Profile Management:** Basic UI to update email, change password, and manage subscription tiers (if applicable).
* **Developer/API Keys:**
  * **UX Security:** When a user generates a new Bearer Token, it is displayed *only once*. The UI must clearly state: "Copy this now, you won't be able to see it again."
  * **Management:** A list of active API keys with the ability to "Revoke" them instantly via a toggle or delete button.

## 6. Global UX Considerations
* **Mobile Responsiveness:** The entire dashboard, especially the data tables and charts, must stack and scroll horizontally gracefully on mobile devices.
* **Dark Mode:** A seamless toggle in the navigation bar that instantly switches the UI theme, mapping all surface colors to low-glare alternatives.
* **Error Handling:** Standardized toast notifications in the bottom right corner for network failures, validation errors, or expired links.
