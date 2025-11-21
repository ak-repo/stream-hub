export default function HomePage() {


  const activities = [
    {
      id: 1,
      type: "file",
      title: "Uploaded report.pdf",
      description: "2 hours ago",
      icon: "📁",
      user: null,
      bgColor: "bg-blue-50",
      iconColor: "text-blue-600",
      borderColor: "border-l-blue-500"
    },
    {
      id: 2,
      type: "message",
      title: "New message from Shibil",
      description: "4 hours ago",
      icon: "💬",
      user: "Shibil",
      bgColor: "bg-teal-50",
      iconColor: "text-teal-600",
      borderColor: "border-l-teal-500"
    },
    {
      id: 3,
      type: "notification",
      title: "You have 3 new notifications",
      description: "Yesterday",
      icon: "🔔",
      user: null,
      bgColor: "bg-red-50",
      iconColor: "text-red-600",
      borderColor: "border-l-red-500"
    },
    {
      id: 4,
      type: "file",
      title: "Shared project-folder with team",
      description: "2 days ago",
      icon: "👥",
      user: "Team",
      bgColor: "bg-purple-50",
      iconColor: "text-purple-600",
      borderColor: "border-l-purple-500"
    }
  ];



  return (
    <div className="w-full h-full space-y-8">
  

   

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
        {/* Recent Activity */}
        <div className="bg-white rounded-2xl p-6 shadow-lg border border-slate-200">
          <div className="flex items-center justify-between mb-6">
            <div>
              <h2 className="text-xl font-semibold text-slate-900">Recent Activity</h2>
              <p className="text-slate-500 text-sm mt-1">
                Your latest file uploads and interactions
              </p>
            </div>
            <button className="text-blue-600 hover:text-blue-700 text-sm font-medium transition-colors">
              View All →
            </button>
          </div>

          <div className="space-y-4">
            {activities.map((activity) => (
              <div
                key={activity.id}
                className={`flex items-center gap-4 p-4 rounded-xl border-l-4 ${activity.borderColor} ${activity.bgColor} hover:shadow-md transition-all duration-200`}
              >
                <div className={`w-12 h-12 ${activity.bgColor} rounded-xl flex items-center justify-center`}>
                  <span className={`text-xl ${activity.iconColor}`}>{activity.icon}</span>
                </div>
                <div className="flex-1">
                  <p className="font-medium text-slate-800">{activity.title}</p>
                  <p className="text-xs text-slate-500 mt-1">{activity.description}</p>
                </div>
                {activity.user && (
                  <div className="px-3 py-1 bg-white rounded-full border border-slate-200">
                    <span className="text-xs font-medium text-slate-700">{activity.user}</span>
                  </div>
                )}
              </div>
            ))}
          </div>
        </div>

        {/* Storage Overview */}
        <div className="bg-white rounded-2xl p-6 shadow-lg border border-slate-200">
          <div className="flex items-center justify-between mb-6">
            <div>
              <h2 className="text-xl font-semibold text-slate-900">Storage Overview</h2>
              <p className="text-slate-500 text-sm mt-1">
                Your cloud storage usage
              </p>
            </div>
            <span className="text-sm font-medium text-slate-700">7.8 GB / 15 GB</span>
          </div>

          {/* Storage Progress */}
          <div className="mb-6">
            <div className="w-full bg-slate-200 rounded-full h-3 mb-2">
              <div 
                className="h-3 rounded-full bg-gradient-to-r from-blue-500 to-teal-500"
                style={{ width: '52%' }}
              ></div>
            </div>
            <div className="flex justify-between text-sm text-slate-500">
              <span>52% used</span>
              <span>7.2 GB free</span>
            </div>
          </div>

          {/* Storage Breakdown */}
          <div className="space-y-4">
            <div className="flex items-center justify-between p-3 bg-slate-50 rounded-lg">
              <div className="flex items-center space-x-3">
                <div className="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center">
                  <span className="text-blue-600">📁</span>
                </div>
                <span className="text-slate-700">Documents</span>
              </div>
              <span className="text-slate-600 font-medium">3.2 GB</span>
            </div>
            <div className="flex items-center justify-between p-3 bg-slate-50 rounded-lg">
              <div className="flex items-center space-x-3">
                <div className="w-8 h-8 bg-red-100 rounded-lg flex items-center justify-center">
                  <span className="text-red-600">🖼️</span>
                </div>
                <span className="text-slate-700">Images</span>
              </div>
              <span className="text-slate-600 font-medium">2.1 GB</span>
            </div>
            <div className="flex items-center justify-between p-3 bg-slate-50 rounded-lg">
              <div className="flex items-center space-x-3">
                <div className="w-8 h-8 bg-purple-100 rounded-lg flex items-center justify-center">
                  <span className="text-purple-600">🎬</span>
                </div>
                <span className="text-slate-700">Videos</span>
              </div>
              <span className="text-slate-600 font-medium">1.8 GB</span>
            </div>
          </div>
        </div>
      </div>

      {/* Recent Files Preview */}
      <div className="bg-white rounded-2xl p-6 shadow-lg border border-slate-200">
        <div className="flex items-center justify-between mb-6">
          <div>
            <h2 className="text-xl font-semibold text-slate-900">Recent Files</h2>
            <p className="text-slate-500 text-sm mt-1">
              Files you've recently worked with
            </p>
          </div>
          <button className="text-blue-600 hover:text-blue-700 text-sm font-medium transition-colors">
            View All Files →
          </button>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
          {['report.pdf', 'presentation.pptx', 'image.jpg', 'data.csv'].map((file, index) => (
            <div key={index} className="p-4 border border-slate-200 rounded-xl hover:shadow-md transition-all duration-200">
              <div className="w-12 h-12 bg-blue-50 rounded-lg flex items-center justify-center mb-3">
                <span className="text-blue-600 text-xl">
                  {file.includes('.pdf') ? '📄' : file.includes('.pptx') ? '📊' : file.includes('.jpg') ? '🖼️' : '📋'}
                </span>
              </div>
              <h3 className="font-medium text-slate-800 truncate">{file}</h3>
              <p className="text-xs text-slate-500 mt-1">Modified 2 hours ago</p>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}