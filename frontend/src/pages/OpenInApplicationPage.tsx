import React, { useState } from 'react';

const OpenInApplicationPage: React.FC = () => {
  const [copied, setCopied] = useState(false);
  const appUrl = 'wails-demo://open';

  const handleCopyUrl = () => {
    navigator.clipboard.writeText(appUrl);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  const handleOpenInBrowser = () => {
    // Open the app URL in the default browser
    // This will trigger the custom protocol handler
    window.open(appUrl, '_blank');
  };

  return (
    <div className="min-h-screen bg-gray-900 p-6">
      <div className="max-w-4xl mx-auto">
        {/* Header */}
        <div className="mb-8">
          <h1 className="text-3xl font-bold text-white mb-2">Open in Application</h1>
          <p className="text-gray-400">
            Open this application from your web browser
          </p>
        </div>

        {/* Main Content */}
        <div className="flex flex-col justify-center items-center min-h-[400px] gap-8">
          {/* Info Card */}
          <div className="w-full max-w-2xl bg-gray-800 rounded-lg p-6 border border-gray-700">
            <div className="mb-6">
              <h2 className="text-xl font-semibold text-white mb-3">How it works</h2>
              <p className="text-gray-300 mb-4">
                You can launch this application directly from your web browser by clicking a link or button.
                The application uses a custom URL scheme to open from external sources.
              </p>
            </div>

            {/* URL Display */}
            <div className="mb-6">
              <label className="block text-sm font-medium text-gray-300 mb-2">
                Application URL Scheme
              </label>
              <div className="flex gap-2">
                <input
                  type="text"
                  value={appUrl}
                  readOnly
                  className="flex-1 px-4 py-3 bg-gray-700 border border-gray-600 rounded-lg text-white font-mono text-sm focus:outline-none focus:ring-2 focus:ring-purple-500"
                />
                <button
                  onClick={handleCopyUrl}
                  className="px-4 py-3 bg-gray-700 text-white rounded-lg hover:bg-gray-600 transition-colors flex items-center gap-2 whitespace-nowrap"
                >
                  {copied ? (
                    <>
                      <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 13l4 4L19 7" />
                      </svg>
                      Copied!
                    </>
                  ) : (
                    <>
                      <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
                      </svg>
                      Copy
                    </>
                  )}
                </button>
              </div>
              <p className="mt-2 text-sm text-gray-400">
                Use this URL in your web pages or applications to open this app
              </p>
            </div>

            {/* Instructions */}
            <div className="bg-gray-700/50 rounded-lg p-4 mb-6">
              <h3 className="text-sm font-semibold text-white mb-2">Usage Example:</h3>
              <code className="block text-xs text-gray-300 font-mono bg-gray-900 p-3 rounded border border-gray-600">
                {'<a href="wails-demo://open">Open Application</a>'}
              </code>
            </div>

            {/* Open Button */}
            <div className="flex justify-center">
              <button
                onClick={handleOpenInBrowser}
                className="px-8 py-4 rounded-lg transition-all flex items-center gap-3 text-lg font-medium shadow-lg transform bg-purple-600 text-white hover:bg-purple-700 hover:shadow-xl hover:scale-105"
              >
                <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
                </svg>
                Open in Application
              </button>
            </div>
          </div>

          {/* Additional Info */}
          <div className="w-full max-w-2xl bg-blue-900/20 border border-blue-700/50 rounded-lg p-4">
            <div className="flex items-start gap-3">
              <svg className="w-5 h-5 text-blue-400 flex-shrink-0 mt-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              <div>
                <p className="text-sm text-blue-300">
                  <strong className="text-blue-200">Note:</strong> The custom URL scheme must be registered during installation. 
                  If clicking the button doesn't open the application, make sure the app is properly installed and the URL scheme is registered.
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default OpenInApplicationPage;

