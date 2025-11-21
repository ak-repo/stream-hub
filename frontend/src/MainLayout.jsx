import { Routes, Route, Navigate, NavLink, Outlet } from "react-router-dom";
import { useState } from "react";

import AuthPage from "./pages/user/AuthPage";
// Dashboard Pages
import HomePage from "./pages/user/HomePage";
import FilesPage from "./pages/user/FilesPage";
import ChatPage from "./pages/user/ChatPage";
import NotificationsPage from "./pages/user/NotificationsPage";
import ProfilePage from "./pages/user/ProfilePage";
import NotFoundPage from "./pages/404";
import Logo from "./components/Logo";

// Protected Route
export function ProtectedRoute({ children }) {
  const token = localStorage.getItem("token");
  return token ? children : <Navigate to="/" replace />;
}

// ----------------------
// Dashboard Layout
// ----------------------
export function DashboardLayout() {
  const [isSidebarOpen, setIsSidebarOpen] = useState(true);

  const navigationItems = [
    { path: "/hub/home", label: "Home", icon: "🏠", badge: null },
    { path: "/hub/files", label: "Files", icon: "📁", badge: "3" },
    { path: "/hub/chat", label: "Chat", icon: "💬", badge: "12" },
    {
      path: "/hub/notifications",
      label: "Notifications",
      icon: "🔔",
      badge: "5",
    },
    { path: "/hub/profile", label: "Profile", icon: "👤", badge: null },
  ];

  return (
    <div className="flex h-screen bg-slate-50">
      {/* Sidebar */}
      <aside
        className={`${
          isSidebarOpen ? "w-64" : "w-20"
        } bg-slate-900 text-white transition-all duration-300 flex flex-col`}
      >
        {/* Logo Section */}
        <div className="p-6 border-b border-slate-700">
          {isSidebarOpen ? (
            <Logo />
          ) : (
            <div className="flex justify-center">
              <div className="w-10 h-10 bg-gradient-to-br from-blue-500 to-teal-400 rounded-xl flex items-center justify-center shadow-lg">
                <svg
                  className="w-6 h-6 text-white"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M9 19V6l12-3v13M9 19c0 1.105-1.343 2-3 2s-3-.895-3-2 1.343-2 3-2 3 .895 3 2zm12-3c0 1.105-1.343 2-3 2s-3-.895-3-2 1.343-2 3-2 3 .895 3 2zM9 10l12-3"
                  />
                </svg>
              </div>
            </div>
          )}
        </div>

        {/* Navigation */}
        <nav className="flex-1 p-4 space-y-2">
          {navigationItems.map((item) => (
            <NavLink
              key={item.path}
              to={item.path}
              className={({ isActive }) =>
                `flex items-center space-x-3 p-3 rounded-xl transition-all duration-200 group ${
                  isActive
                    ? "bg-blue-500 text-white shadow-lg shadow-blue-500/25"
                    : "text-slate-300 hover:bg-slate-800 hover:text-white"
                }`
              }
            >
              <span className="text-xl flex-shrink-0">{item.icon}</span>
              {isSidebarOpen && (
                <>
                  <span className="font-medium flex-1">{item.label}</span>
                  {item.badge && (
                    <span className="bg-red-500 text-white text-xs px-2 py-1 rounded-full">
                      {item.badge}
                    </span>
                  )}
                </>
              )}
            </NavLink>
          ))}
        </nav>

        {/* User Profile & Toggle */}
        <div className="p-4 border-t border-slate-700">
          <div className="flex items-center space-x-3 p-3 rounded-xl bg-slate-800">
            <div className="w-8 h-8 bg-teal-500 rounded-full flex items-center justify-center text-white font-semibold text-sm">
              JD
            </div>
            {isSidebarOpen && (
              <div className="flex-1 min-w-0">
                <p className="text-white font-medium text-sm truncate">
                  John Doe
                </p>
                <p className="text-slate-400 text-xs truncate">
                  john@streamhub.com
                </p>
              </div>
            )}
          </div>

          {/* Toggle Sidebar Button */}
          <button
            onClick={() => setIsSidebarOpen(!isSidebarOpen)}
            className="w-full mt-3 p-2 text-slate-400 hover:text-white hover:bg-slate-800 rounded-lg transition-colors"
          >
            {isSidebarOpen ? "← Collapse" : "→ Expand"}
          </button>
        </div>
      </aside>

      {/* Main Content */}
      <main className="flex-1 flex flex-col overflow-hidden">
    

        {/* Page Content */}
        <div className="flex-1 overflow-auto p-6">
          <div className=" mx-auto">
            <Outlet />
          </div>
        </div>
      </main>
    </div>
  );
}

// Main Routes
export default function AppRoutes() {
  return (
    <Routes>
      {/* Authentication Page */}
      <Route path="/" element={<AuthPage />} />
      <Route path="*" element={<NotFoundPage />} />

      {/* Dashboard Routes (Protected) */}
      <Route
        path="/hub"
        element={
          <ProtectedRoute>
            <DashboardLayout />
          </ProtectedRoute>
        }
      >
        <Route path="home" element={<HomePage />} />
        <Route path="files" element={<FilesPage />} />
        <Route path="chat" element={<ChatPage />} />
        <Route path="notifications" element={<NotificationsPage />} />
        <Route path="profile" element={<ProfilePage />} />

        {/* Default redirect to /hub/home */}
        <Route index element={<Navigate to="/hub/home" replace />} />
      </Route>
    </Routes>
  );
}
