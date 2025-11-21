import { useState } from "react";
import { useAuth } from "../../context/context";
import { useNavigate } from "react-router-dom";

export default function AuthPage() {
  const [isLogin, setIsLogin] = useState(true);
  const navigate = useNavigate();
  const [form, setForm] = useState({ username: "", password: "", email: "" });
  const { login, register, isAuthenticated } = useAuth();

  const handleSubmit = async (e) => {
    e.preventDefault();
    console.log("form: ", form);
    if (isLogin) {
      await login(form.email, form.password);
    } else {
      await register(form);
    }
    if (isAuthenticated) {
      navigate("/hub");
    }
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setForm((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  // OAuth providers data
  const oauthProviders = [
    { name: "Google", icon: "🔍", color: "hover:bg-red-50 border-red-200" },
    { name: "GitHub", icon: "💻", color: "hover:bg-gray-50 border-gray-200" },
    {
      name: "Microsoft",
      icon: "🔷",
      color: "hover:bg-blue-50 border-blue-200",
    },
    { name: "Apple", icon: "🍎", color: "hover:bg-black border-gray-300" },
  ];

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-50 via-blue-50 to-teal-50 flex">
      {/* Left Side - Branding */}
      <div className="hidden lg:flex lg:flex-1 lg:flex-col lg:justify-center lg:px-12 lg:py-12 bg-gradient-to-br from-blue-600 to-teal-500 text-white">
        <div className="mx-auto max-w-md">
          <div className="flex items-center space-x-3 mb-8">
            <div className="w-12 h-12 bg-white/20 rounded-xl flex items-center justify-center backdrop-blur-sm">
              <span className="text-2xl font-bold">📊</span>
            </div>
            <h1 className="text-3xl font-bold">StreamHub</h1>
          </div>

          <h2 className="text-4xl font-bold mb-6 leading-tight">
            Your Central Hub for{" "}
            <span className="text-teal-200">Files, Chat & Collaboration</span>
          </h2>

          <p className="text-blue-100 text-lg mb-8 leading-relaxed">
            Streamline your workflow with seamless file sharing, real-time
            communication, and smart notifications—all in one powerful platform.
          </p>

          <div className="space-y-4">
            <div className="flex items-center space-x-3">
              <div className="w-8 h-8 bg-teal-400 rounded-full flex items-center justify-center">
                <span className="text-white text-sm">✓</span>
              </div>
              <span className="text-blue-100">
                Secure file storage & sharing
              </span>
            </div>
            <div className="flex items-center space-x-3">
              <div className="w-8 h-8 bg-teal-400 rounded-full flex items-center justify-center">
                <span className="text-white text-sm">✓</span>
              </div>
              <span className="text-blue-100">
                Real-time team collaboration
              </span>
            </div>
            <div className="flex items-center space-x-3">
              <div className="w-8 h-8 bg-teal-400 rounded-full flex items-center justify-center">
                <span className="text-white text-sm">✓</span>
              </div>
              <span className="text-blue-100">Smart activity tracking</span>
            </div>
          </div>
        </div>
      </div>

      {/* Right Side - Login Form */}
      <div className="flex-1 flex flex-col justify-center px-6 py-12 lg:px-12">
        <div className="mx-auto w-full max-w-md">
          {/* Mobile Logo */}
          <div className="lg:hidden flex justify-center mb-8">
            <div className="flex items-center space-x-3">
              <div className="w-12 h-12 bg-blue-600 rounded-xl flex items-center justify-center">
                <span className="text-2xl font-bold text-white">📊</span>
              </div>
              <h1 className="text-2xl font-bold text-slate-800">StreamHub</h1>
            </div>
          </div>

          <div className="text-center lg:text-left mb-8">
            <h2 className="text-3xl font-bold text-slate-800">
              {isLogin ? "Welcome back" : "Join StreamHub"}
            </h2>
            <p className="mt-2 text-slate-600">
              {isLogin
                ? "Sign in to continue your workflow"
                : "Create your account to get started"}
            </p>
          </div>

          {/* Toggle Switch */}
          <div className="mb-8">
            <div className="bg-slate-100 rounded-2xl p-1.5 shadow-inner">
              <div className="grid grid-cols-2 gap-1">
                <button
                  onClick={() => setIsLogin(true)}
                  className={`py-3 px-6 rounded-xl text-sm font-semibold transition-all duration-200 ${
                    isLogin
                      ? "bg-white text-blue-600 shadow-sm"
                      : "text-slate-600 hover:text-slate-800"
                  }`}
                >
                  Sign In
                </button>
                <button
                  onClick={() => setIsLogin(false)}
                  className={`py-3 px-6 rounded-xl text-sm font-semibold transition-all duration-200 ${
                    !isLogin
                      ? "bg-white text-blue-600 shadow-sm"
                      : "text-slate-600 hover:text-slate-800"
                  }`}
                >
                  Sign Up
                </button>
              </div>
            </div>
          </div>

          <div className="bg-white py-8 px-8 shadow-xl rounded-2xl border border-slate-200">
            {/* OAuth Section */}
            <div className="mb-8">
              <div className="grid grid-cols-2 gap-3">
                {oauthProviders.map((provider) => (
                  <button
                    key={provider.name}
                    className={`flex items-center justify-center space-x-2 py-3 px-4 rounded-xl border transition-all duration-200 ${provider.color} hover:shadow-sm group`}
                  >
                    <span className="text-lg">{provider.icon}</span>
                    <span className="text-sm font-medium text-slate-700 group-hover:text-slate-900">
                      {provider.name}
                    </span>
                  </button>
                ))}
              </div>

              <div className="relative mt-6">
                <div className="absolute inset-0 flex items-center">
                  <div className="w-full border-t border-slate-200"></div>
                </div>
                <div className="relative flex justify-center text-sm">
                  <span className="px-2 bg-white text-slate-500">
                    Or continue with email
                  </span>
                </div>
              </div>
            </div>

            <form className="space-y-5">
              {!isLogin && (
                <div>
                  <label
                    htmlFor="username"
                    className="block text-sm font-medium text-slate-700 mb-2"
                  >
                    Full Name
                  </label>
                  <div className="mt-1">
                    <input
                      id="username"
                      name="username"
                      type="text"
                      value={form.username}
                      onChange={(e) => handleChange(e)}
                      required
                      className="block w-full rounded-xl border border-slate-300 px-4 py-3 text-slate-900 placeholder-slate-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 bg-slate-50 hover:bg-white"
                      placeholder="Enter your full name"
                    />
                  </div>
                </div>
              )}

              <div>
                <label
                  htmlFor="email"
                  className="block text-sm font-medium text-slate-700 mb-2"
                >
                  Email address
                </label>
                <div className="mt-1">
                  <input
                    id="email"
                    name="email"
                    type="email"
                    value={form.email}
                    onChange={(e) => handleChange(e)}
                    required
                    autoComplete="email"
                    className="block w-full rounded-xl border border-slate-300 px-4 py-3 text-slate-900 placeholder-slate-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 bg-slate-50 hover:bg-white"
                    placeholder="Enter your email"
                  />
                </div>
              </div>

              <div>
                <div className="flex items-center justify-between mb-2">
                  <label
                    htmlFor="password"
                    className="block text-sm font-medium text-slate-700"
                  >
                    Password
                  </label>
                  {isLogin && (
                    <div className="text-sm">
                      <a
                        href="#"
                        className="font-medium text-blue-600 hover:text-blue-500 transition-colors"
                      >
                        Forgot password?
                      </a>
                    </div>
                  )}
                </div>
                <div className="mt-1">
                  <input
                    onChange={(e) => handleChange(e)}
                    id="password"
                    name="password"
                    type="password"
                    value={form.password}
                    required
                    autoComplete={isLogin ? "current-password" : "new-password"}
                    className="block w-full rounded-xl border border-slate-300 px-4 py-3 text-slate-900 placeholder-slate-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 bg-slate-50 hover:bg-white"
                    placeholder={
                      isLogin ? "Enter your password" : "Create a password"
                    }
                  />
                </div>
              </div>

              {!isLogin && (
                <div className="flex items-center space-x-2">
                  <input
                    type="checkbox"
                    id="terms"
                    className="w-4 h-4 text-blue-600 border-slate-300 rounded focus:ring-blue-500"
                  />
                  <label htmlFor="terms" className="text-sm text-slate-600">
                    I agree to the{" "}
                    <a
                      href="#"
                      className="text-blue-600 hover:text-blue-500 font-medium"
                    >
                      Terms of Service
                    </a>{" "}
                    and{" "}
                    <a
                      href="#"
                      className="text-blue-600 hover:text-blue-500 font-medium"
                    >
                      Privacy Policy
                    </a>
                  </label>
                </div>
              )}

              <div>
                <button
                  onClick={(e) => handleSubmit(e)}
                  type="submit"
                  className="flex w-full justify-center rounded-xl bg-blue-600 py-3.5 px-4 text-sm font-semibold text-white hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition-all duration-200 shadow-lg hover:shadow-xl"
                >
                  {isLogin
                    ? "Sign in to StreamHub"
                    : "Create StreamHub account"}
                </button>
              </div>
            </form>

            <div className="mt-6 text-center">
              <p className="text-sm text-slate-600">
                {isLogin
                  ? "Don't have an account?"
                  : "Already have an account?"}{" "}
                <button
                  onClick={() => setIsLogin(!isLogin)}
                  className="font-medium text-blue-600 hover:text-blue-500 transition-colors"
                >
                  {isLogin ? "Sign up" : "Sign in"}
                </button>
              </p>
            </div>
          </div>

          {/* Footer */}
          <div className="mt-8 text-center">
            <p className="text-xs text-slate-500">
              © 2024 StreamHub. All rights reserved.
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}
