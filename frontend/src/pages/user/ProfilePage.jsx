export default function ProfilePage() {
  return (
    <div>
      <h1 className="text-2xl font-semibold mb-4">Profile</h1>

      <div className="bg-white p-6 rounded-xl shadow-sm w-full max-w-lg">
        {/* User Avatar */}
        <div className="flex items-center gap-4 mb-6">
          <div className="w-20 h-20 bg-gray-300 rounded-full"></div>
          <div>
            <p className="text-xl font-semibold">Ananda Krishnan</p>
            <p className="text-gray-500 text-sm">streamhub@user.com</p>
          </div>
        </div>

        {/* Info */}
        <div className="space-y-4">
          <div>
            <p className="text-gray-500">Username</p>
            <input
              type="text"
              value="ananda"
              className="border p-2 rounded-xl w-full"
            />
          </div>

          <div>
            <p className="text-gray-500">Email</p>
            <input
              type="email"
              value="ananda@streamhub.com"
              className="border p-2 rounded-xl w-full"
            />
          </div>

          <button className="w-full bg-blue-600 text-white p-2 rounded-xl mt-4">
            Save Changes
          </button>
        </div>
      </div>
    </div>
  );
}
