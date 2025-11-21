export default function ChatPage() {
  const messages = [
    {
      id: 1,
      sender: "You",
      text: "Hello! 👋",
      time: "2:30 PM",
      avatar: "👤",
      isYou: true,
    },
    {
      id: 2,
      sender: "Shibil",
      text: "Hey bro, how is StreamHub going?",
      time: "2:31 PM",
      avatar: "🚀",
      isYou: false,
    },
    {
      id: 3,
      sender: "You",
      text: "Working on UI now!",
      time: "2:32 PM",
      avatar: "👤",
      isYou: true,
    },
  ];

  const conversations = [
    {
      id: 1,
      name: "Shibil",
      avatar: "🚀",
      lastMessage: "Hey bro, how is StreamHub...",
      time: "2:31 PM",
      unread: 1,
      active: true,
    },
    {
      id: 2,
      name: "Design Team",
      avatar: "🎨",
      lastMessage: "New mockups ready",
      time: "1:15 PM",
      unread: 0,
      active: false,
    },
    {
      id: 3,
      name: "Sarah Chen",
      avatar: "🌟",
      lastMessage: "Thanks for the feedback!",
      time: "12:45 PM",
      unread: 0,
      active: false,
    },
  ];

  return (
    <div className="w-full h-full flex flex-col">
      {/* Header */}
      <div className="mb-6">
        <h1 className="text-3xl font-bold text-slate-900">Chat</h1>
        <p className="text-slate-600 mt-2">
          Message your team and collaborators
        </p>
      </div>

      <div className="flex-1 flex gap-6 min-h-0">
        {/* Conversations Sidebar - Compact */}
        <div className="w-80 bg-white rounded-2xl shadow-sm border border-slate-200 flex flex-col">
          <div className="p-4 border-b border-slate-200">
            <div className="relative">
              <input
                type="text"
                placeholder="Search conversations..."
                className="w-full bg-slate-50 border border-slate-300 rounded-xl px-4 py-2.5 pl-10 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent text-sm"
              />
              <span className="absolute left-3 top-2.5 text-slate-400 text-sm">
                🔍
              </span>
            </div>
          </div>

          <div className="flex-1 overflow-y-auto">
            {conversations.map((convo) => (
              <div
                key={convo.id}
                className={`flex items-center gap-3 p-3 border-b border-slate-100 cursor-pointer transition-colors ${
                  convo.active
                    ? "bg-blue-50 border-l-4 border-l-blue-500"
                    : "hover:bg-slate-50"
                }`}
              >
                <div className="w-10 h-10 bg-gradient-to-br from-blue-500 to-teal-400 rounded-xl flex items-center justify-center text-white font-semibold text-sm">
                  {convo.avatar}
                </div>
                <div className="flex-1 min-w-0">
                  <div className="flex items-center justify-between mb-1">
                    <h3 className="font-semibold text-slate-900 text-sm truncate">
                      {convo.name}
                    </h3>
                    <span className="text-xs text-slate-500">{convo.time}</span>
                  </div>
                  <p className="text-sm text-slate-600 truncate">
                    {convo.lastMessage}
                  </p>
                </div>
                {convo.unread > 0 && (
                  <div className="w-5 h-5 bg-red-500 rounded-full flex items-center justify-center">
                    <span className="text-white text-xs font-semibold">
                      {convo.unread}
                    </span>
                  </div>
                )}
              </div>
            ))}
          </div>
        </div>

        {/* Chat Area */}
        <div className="flex-1 flex flex-col bg-white rounded-2xl shadow-sm border border-slate-200 min-h-0">
          {/* Chat Header */}
          <div className="p-4 border-b border-slate-200">
            <div className="flex items-center gap-3">
              <div className="w-8 h-8 bg-gradient-to-br from-blue-500 to-teal-400 rounded-xl flex items-center justify-center text-white font-semibold text-sm">
                🚀
              </div>
              <div>
                <h2 className="font-semibold text-slate-900 text-sm">Shibil</h2>
                <p className="text-xs text-slate-500">Online</p>
              </div>
            </div>
          </div>

          {/* Messages Container */}
          <div className="flex-1 overflow-y-auto p-4 bg-slate-50">
            <div className="space-y-3">
              {messages.map((msg) => (
                <div
                  key={msg.id}
                  className={`flex gap-2 ${
                    msg.isYou ? "flex-row-reverse" : ""
                  }`}
                >
                  {/* Avatar - Smaller */}
                  <div
                    className={`w-6 h-6 rounded-full flex items-center justify-center flex-shrink-0 text-xs ${
                      msg.isYou
                        ? "bg-blue-500 text-white"
                        : "bg-teal-500 text-white"
                    }`}
                  >
                    {msg.avatar}
                  </div>

                  {/* Message Bubble */}
                  <div
                    className={`max-w-[75%] ${msg.isYou ? "text-right" : ""}`}
                  >
                    <div
                      className={`inline-block p-3 rounded-2xl ${
                        msg.isYou
                          ? "bg-blue-500 text-white rounded-br-none"
                          : "bg-white text-slate-900 border border-slate-200 rounded-bl-none shadow-sm"
                      }`}
                    >
                      <p className="text-xs mb-1 opacity-90">{msg.sender}</p>
                      <p className="text-sm leading-relaxed">{msg.text}</p>
                    </div>
                    <p
                      className={`text-xs text-slate-500 mt-1 ${
                        msg.isYou ? "text-right" : ""
                      }`}
                    >
                      {msg.time}
                    </p>
                  </div>
                </div>
              ))}
            </div>
          </div>

          {/* Input Area */}
          <div className="p-4 border-t border-slate-200">
            <div className="flex gap-2">
              <div className="flex-1 bg-slate-50 rounded-xl border border-slate-300 focus-within:border-blue-500 focus-within:ring-2 focus-within:ring-blue-200 transition-all">
                <input
                  type="text"
                  placeholder="Type your message..."
                  className="w-full bg-transparent border-none focus:outline-none px-3 py-2 text-slate-900 placeholder-slate-500 text-sm"
                />
              </div>
              <button className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-xl font-medium transition-colors shadow-sm hover:shadow-md flex items-center gap-1 text-sm">
                <span>Send</span>
                <span className="text-xs">↑</span>
              </button>
            </div>

            {/* Quick Actions - Compact */}
            <div className="flex items-center gap-2 mt-2">
              <button className="text-slate-500 hover:text-slate-700 p-1 rounded-lg hover:bg-slate-100 transition-colors text-sm">
                📎
              </button>
              <button className="text-slate-500 hover:text-slate-700 p-1 rounded-lg hover:bg-slate-100 transition-colors text-sm">
                🖼️
              </button>
              <button className="text-slate-500 hover:text-slate-700 p-1 rounded-lg hover:bg-slate-100 transition-colors text-sm">
                😊
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
