# Frontend Service (Next.js)

The frontend is a modern, high-performance web application built with **Next.js 14** using the **App Router** and **Tailwind CSS**.

## 🚀 Features

- **Dynamic Analytics Dashboard**: Visualize link performance with real-time click tracking and trend analysis.
- **Link Management**: Easily create, edit, and delete short links with a clean, responsive UI.
- **Authentication**: Seamless JWT-based login and registration flows.
- **Responsive Design**: Mobile-first design using **Tailwind CSS**.

---

## 🏗️ Architecture

- **Framework**: Next.js 14 (App Router)
- **Styling**: Tailwind CSS for rapid UI development.
- **State Management**: Built-in React hooks and Next.js server components for optimized data fetching.
- **Data Fetching**: Communicates with the backend API via standard Fetch API with built-in JWT handling.

### Key Directories:

- `app/`: Next.js App Router structure (pages, layouts, and components).
- `components/`: Reusable React components (buttons, charts, cards).
- `lib/`: Utility functions and API client wrappers.
- `public/`: Static assets (images, icons).

---

## 💻 Local Development

Before starting, ensure the backend services are running.

### 1. Install Dependencies
```bash
npm install
```

### 2. Configure Environment
Create a `.env.local` if you need to override the default API URL:
```bash
NEXT_PUBLIC_API_URL=http://localhost:5010
```

### 3. Start Development Server
```bash
npm run dev
```

The app will be available at [http://localhost:3000](http://localhost:3000).

---

## 🛠️ Build and Production

To build the optimized production bundle:

```bash
npm run build
```

The output will be in the `.next` directory. For static exports (if configured), see the `out` directory.

---

## 🎨 UI & Styling

- **Icons**: [Lucide React](https://lucide.dev/)
- **Charts**: [Recharts](https://recharts.org/) for data visualization.
- **Fonts**: Geist Mono & Sans for a modern, sleek aesthetic.
