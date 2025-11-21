export default function NotificationsPage() {
  const notifications = [
    { id: 1, text: "Your file 'video.mp4' has finished uploading.", time: "2 min ago" },
    { id: 2, text: "New login detected on your account.", time: "1 hr ago" },
    { id: 3, text: "Shibil sent you a message.", time: "Yesterday" },
  ];

  return (
    <div>
      <h1 className="text-2xl font-semibold mb-4">Notifications</h1>

      <div className="bg-white p-4 rounded-xl shadow-sm">
        {notifications.map((n) => (
          <div
            key={n.id}
            className="border-b p-3 last:border-none hover:bg-gray-50 rounded"
          >
            <p className="font-medium">{n.text}</p>
            <p className="text-sm text-gray-500">{n.time}</p>
          </div>
        ))}
      </div>
    </div>
  );
}
