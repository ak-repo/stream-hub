import { useNavigate } from "react-router-dom";
import Logo from "../components/Logo";
import { useAuth } from "../context/context";

function Navbar() {
  const navigate = useNavigate();
  const { isAuthenticated } = useAuth();
  return (
    <div>
      <nav className="bg-emerald-500 shadow-lg border-b border-emerald-600">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <Logo />
            <div className="hidden md:flex items-center space-x-8">
              <p
                onClick={() => navigate("/")}
                className="text-white hover:text-emerald-100 font-medium transition-colors duration-200 cursor-pointer"
              >
                Home
              </p>
              <p
                onClick={() => navigate("/products")}
                className="text-white hover:text-emerald-100 font-medium transition-colors duration-200 cursor-pointer"
              >
                Products
              </p>
              <p
                onClick={() => navigate("/about")}
                className="text-white hover:text-emerald-100 font-medium transition-colors duration-200 cursor-pointer"
              >
                About
              </p>
              <p
                onClick={() => navigate("/contact")}
                className="text-white hover:text-emerald-100 font-medium transition-colors duration-200 cursor-pointer"
              >
                Contact
              </p>
            </div>
            <div className="flex items-center space-x-4">
              {isAuthenticated ? (
                <>
                  <button
                    className="text-white hover:text-emerald-100 font-medium transition-colors duration-200 bg-emerald-600 hover:bg-emerald-700 px-4 py-2 rounded-lg"
                    onClick={() => navigate("/cart")}
                  >
                    Cart
                  </button>
                  <button
                    className="text-white hover:text-emerald-100 font-medium transition-colors duration-200 bg-emerald-600 hover:bg-emerald-700 px-4 py-2 rounded-lg"
                    onClick={() => navigate("/wishlist")}
                  >
                    Wishlist
                  </button>
                  <button
                    className="text-white hover:text-emerald-100 font-medium transition-colors duration-200 bg-emerald-600 hover:bg-emerald-700 px-4 py-2 rounded-lg"
                    onClick={() => navigate("/profile")}
                  >
                    Profile
                  </button>
                </>
              ) : (
                <button
                  onClick={() => navigate("/auth")}
                  className="text-white hover:text-emerald-100 font-medium transition-colors duration-200 bg-emerald-600 hover:bg-emerald-700 px-6 py-2 rounded-lg"
                >
                  Login
                </button>
              )}
            </div>
          </div>
        </div>
      </nav>
    </div>
  );
}

export default Navbar;
