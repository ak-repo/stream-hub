export default function FilesPage() {
  const files = [
    {
      id: 1,
      name: "project-brief.pdf",
      size: "2.1 MB",
      type: "PDF",
      date: "2 hours ago",
    },
    {
      id: 2,
      name: "presentation-deck.pptx",
      size: "15.4 MB",
      type: "Presentation",
      date: "Yesterday",
    },
    {
      id: 3,
      name: "team-photo.jpg",
      size: "4.2 MB",
      type: "Image",
      date: "3 days ago",
    },
    {
      id: 4,
      name: "quarterly-report.xlsx",
      size: "800 KB",
      type: "Spreadsheet",
      date: "1 week ago",
    },
  ];

  const fileTypeIcons = {
    PDF: "📄",
    Presentation: "📊",
    Image: "🖼️",
    Spreadsheet: "📋",
    Video: "🎬",
    Audio: "🎵",
    Document: "📝",
    Archive: "📦",
  };

  return (
    <div className="min-h-screen bg-slate-50 p-6">
      <div className=" mx-auto">
        {/* Header */}
        <div className="mb-8">
          <h1 className="text-3xl font-bold text-slate-900">Files</h1>
          <p className="text-slate-600 mt-2">
            Manage and organize your documents
          </p>
        </div>

        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          {/* Main Content */}
          <div className="lg:col-span-2 space-y-6">
            {/* Upload Section */}
            <div className="bg-white rounded-2xl shadow-sm border border-slate-200 p-8">
              <div className="text-center">
                <div className="w-16 h-16 bg-blue-50 rounded-2xl flex items-center justify-center mx-auto mb-4">
                  <span className="text-2xl text-blue-600">📤</span>
                </div>
                <h2 className="text-xl font-semibold text-slate-900 mb-2">
                  Upload Files
                </h2>
                <p className="text-slate-600 mb-6">
                  Drag and drop or click to browse
                </p>

                {/* Upload Box */}
                <div className="border-2 border-dashed border-slate-300 rounded-2xl p-8 hover:border-blue-500 transition-colors cursor-pointer bg-slate-50 hover:bg-blue-50">
                  <div className="flex flex-col items-center">
                    <span className="text-4xl mb-3">📁</span>
                    <p className="text-slate-700 font-medium mb-1">
                      Drop files here
                    </p>
                    <p className="text-slate-500 text-sm">
                      or click to select from your device
                    </p>
                    <p className="text-slate-400 text-xs mt-2">
                      Supports PDF, DOC, PPT, JPG, MP4
                    </p>
                  </div>
                </div>

                {/* Upload Button */}
                <button className="mt-6 w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-3 px-6 rounded-xl transition-colors shadow-sm hover:shadow-md">
                  Browse Files
                </button>
              </div>
            </div>

            {/* Quick Actions */}
            <div className="bg-white rounded-2xl shadow-sm border border-slate-200 p-6">
              <h3 className="text-lg font-semibold text-slate-900 mb-4">
                Quick Actions
              </h3>
              <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
                <button className="flex flex-col items-center p-4 bg-slate-50 hover:bg-blue-50 rounded-xl transition-colors group">
                  <span className="text-2xl mb-2 group-hover:scale-110 transition-transform">
                    📁
                  </span>
                  <span className="text-sm font-medium text-slate-700">
                    New Folder
                  </span>
                </button>
                <button className="flex flex-col items-center p-4 bg-slate-50 hover:bg-teal-50 rounded-xl transition-colors group">
                  <span className="text-2xl mb-2 group-hover:scale-110 transition-transform">
                    🔗
                  </span>
                  <span className="text-sm font-medium text-slate-700">
                    Share
                  </span>
                </button>
                <button className="flex flex-col items-center p-4 bg-slate-50 hover:bg-purple-50 rounded-xl transition-colors group">
                  <span className="text-2xl mb-2 group-hover:scale-110 transition-transform">
                    📊
                  </span>
                  <span className="text-sm font-medium text-slate-700">
                    Storage
                  </span>
                </button>
                <button className="flex flex-col items-center p-4 bg-slate-50 hover:bg-red-50 rounded-xl transition-colors group">
                  <span className="text-2xl mb-2 group-hover:scale-110 transition-transform">
                    🗑️
                  </span>
                  <span className="text-sm font-medium text-slate-700">
                    Trash
                  </span>
                </button>
              </div>
            </div>
          </div>

          {/* File List Sidebar */}
          <div className="lg:col-span-1">
            <div className="bg-white rounded-2xl shadow-sm border border-slate-200 p-6 sticky top-6">
              <div className="flex items-center justify-between mb-6">
                <h2 className="text-xl font-semibold text-slate-900">
                  Recent Files
                </h2>
                <span className="text-sm text-blue-600 font-medium cursor-pointer hover:text-blue-700">
                  View All
                </span>
              </div>

              <div className="space-y-4">
                {files.map((file) => (
                  <div
                    key={file.id}
                    className="flex items-center gap-4 p-3 rounded-xl hover:bg-slate-50 transition-colors group cursor-pointer"
                  >
                    <div className="w-12 h-12 bg-blue-50 rounded-xl flex items-center justify-center flex-shrink-0">
                      <span className="text-lg">
                        {fileTypeIcons[file.type] || "📄"}
                      </span>
                    </div>
                    <div className="flex-1 min-w-0">
                      <p className="font-medium text-slate-900 truncate">
                        {file.name}
                      </p>
                      <div className="flex items-center gap-2 text-sm text-slate-500">
                        <span>{file.type}</span>
                        <span>•</span>
                        <span>{file.size}</span>
                      </div>
                      <p className="text-xs text-slate-400 mt-1">{file.date}</p>
                    </div>
                    <button className="opacity-0 group-hover:opacity-100 p-2 hover:bg-slate-200 rounded-lg transition-all">
                      <span className="text-slate-600">⋯</span>
                    </button>
                  </div>
                ))}
              </div>

              {/* Storage Info */}
              <div className="mt-8 pt-6 border-t border-slate-200">
                <div className="flex justify-between items-center mb-2">
                  <span className="text-sm font-medium text-slate-700">
                    Storage
                  </span>
                  <span className="text-sm text-slate-600">7.2 GB / 15 GB</span>
                </div>
                <div className="w-full bg-slate-200 rounded-full h-2">
                  <div
                    className="bg-gradient-to-r from-blue-500 to-teal-500 h-2 rounded-full transition-all duration-500"
                    style={{ width: "48%" }}
                  ></div>
                </div>
                <p className="text-xs text-slate-500 mt-2">
                  48% used • 7.8 GB available
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
