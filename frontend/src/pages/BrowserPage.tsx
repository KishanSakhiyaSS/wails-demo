import React, { useState } from 'react';
import { BrowserOpenURL } from '../../wailsjs/runtime/runtime';

const BrowserPage: React.FC = () => {
  const [url, setUrl] = useState('');
  const [error, setError] = useState<string | null>(null);
  const [touched, setTouched] = useState(false);

  const validateURL = (urlString: string): boolean => {
    // Clear any previous errors first
    setError(null);

    if (!urlString || urlString.trim() === '') {
      setError('URL cannot be empty');
      return false;
    }

    const trimmedUrl = urlString.trim();

    // Check for invalid characters
    if (trimmedUrl.includes(' ')) {
      setError('URL cannot contain spaces');
      return false;
    }

    // Check for basic invalid patterns
    if (trimmedUrl === '.' || trimmedUrl === '..' || trimmedUrl.startsWith('..')) {
      setError('Invalid URL format');
      return false;
    }

    // Try to parse as URL
    try {
      let formattedUrl = trimmedUrl;
      
      // If URL doesn't start with http:// or https://, add https://
      if (!formattedUrl.match(/^https?:\/\//i)) {
        // Check if it looks like it might be a valid domain
        if (!formattedUrl.match(/^[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}/)) {
          setError('Invalid URL format. Please enter a valid URL (e.g., www.example.com or https://example.com)');
          return false;
        }
        formattedUrl = 'https://' + formattedUrl;
      }

      // Validate URL using URL constructor
      const urlObj = new URL(formattedUrl);
      
      // Validate protocol
      if (!['http:', 'https:'].includes(urlObj.protocol)) {
        setError('URL must use http:// or https:// protocol');
        return false;
      }

      // Validate hostname exists and is not empty
      if (!urlObj.hostname || urlObj.hostname.length === 0) {
        setError('Invalid URL: hostname is required');
        return false;
      }

      // Validate hostname doesn't contain invalid characters
      if (urlObj.hostname.includes('..') || urlObj.hostname.startsWith('.') || urlObj.hostname.endsWith('.')) {
        setError('Invalid hostname format');
        return false;
      }

      // Validate hostname has at least one dot (for domain extensions like .com, .org, etc.)
      // But allow localhost and IP addresses
      const isLocalhost = urlObj.hostname === 'localhost';
      const isIPAddress = /^(\d{1,3}\.){3}\d{1,3}$/.test(urlObj.hostname);
      const hasValidDomain = urlObj.hostname.includes('.') || isLocalhost || isIPAddress;
      
      if (!hasValidDomain && urlObj.hostname.length > 0) {
        setError('Invalid URL: hostname must be a valid domain (e.g., example.com)');
        return false;
      }

      // If all validations pass, clear error and return true
      setError(null);
      return true;
    } catch (err) {
      setError('Invalid URL format. Please enter a valid URL (e.g., www.example.com or https://example.com)');
      return false;
    }
  };

  const handleOpenBrowser = () => {
    setTouched(true);
    
    // Validate URL first - this will set error if invalid
    const isValid = validateURL(url);
    
    // Only open browser if validation passes
    if (isValid) {
      // Format URL with https:// if not present
      let formattedUrl = url.trim();
      if (!formattedUrl.match(/^https?:\/\//i)) {
        formattedUrl = 'https://' + formattedUrl;
      }
      
      // Double-check the formatted URL is valid before opening
      try {
        new URL(formattedUrl);
        BrowserOpenURL(formattedUrl);
        // Clear error on success
        setError(null);
      } catch (err) {
        setError('Failed to open URL. Please check the URL format.');
      }
    } else {
      // Validation failed - error is already set by validateURL
      // Don't open browser
      return;
    }
  };

  const handleUrlChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const newUrl = e.target.value;
    setUrl(newUrl);
    // Clear error when user starts typing (they're fixing it)
    if (error && touched) {
      // Only clear if user is actively typing and field was previously touched
      setError(null);
    }
  };

  const handleBlur = () => {
    setTouched(true);
    if (url.trim() === '') {
      setError('URL cannot be empty');
    } else {
      validateURL(url);
    }
  };

  const handleKeyPress = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter') {
      handleOpenBrowser();
    }
  };

  return (
    <div className="min-h-screen bg-gray-900 p-6">
      <div className="max-w-6xl mx-auto">
        {/* Header */}
        <div className="mb-8">
          <h1 className="text-3xl font-bold text-white mb-2">Browser</h1>
          <p className="text-gray-400">
            Open links in your default web browser
          </p>
        </div>

        {/* URL Input and Button */}
        <div className="flex flex-col justify-center items-center min-h-[400px] gap-6">
          <div className="w-full max-w-md">
            <label htmlFor="url-input" className="block text-sm font-medium text-gray-300 mb-2">
              Enter URL
            </label>
            <input
              id="url-input"
              type="text"
              value={url}
              onChange={handleUrlChange}
              onBlur={handleBlur}
              onKeyPress={handleKeyPress}
              placeholder="https://www.example.com or www.example.com"
              className={`w-full px-4 py-3 bg-gray-800 border-2 rounded-lg text-white placeholder-gray-500 focus:outline-none focus:ring-2 transition-colors ${
                error 
                  ? 'border-red-500 focus:ring-red-500 focus:border-red-500' 
                  : 'border-gray-700 focus:ring-purple-500 focus:border-purple-500'
              }`}
            />
            {error && touched && (
              <div className="mt-2 p-3 bg-red-900/30 border border-red-500 rounded-lg">
                <p className="text-sm text-red-400 flex items-center gap-2">
                  <svg className="w-5 h-5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                  <span>{error}</span>
                </p>
              </div>
            )}
          </div>

          <button
            onClick={handleOpenBrowser}
            disabled={!url || url.trim() === '' || (error !== null && touched)}
            className={`px-8 py-4 rounded-lg transition-all flex items-center gap-3 text-lg font-medium shadow-lg transform ${
              !url || url.trim() === '' || (error !== null && touched)
                ? 'bg-gray-600 text-gray-400 cursor-not-allowed'
                : 'bg-purple-600 text-white hover:bg-purple-700 hover:shadow-xl hover:scale-105'
            }`}
          >
            <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
            </svg>
            Open in browser
          </button>
        </div>
      </div>
    </div>
  );
};

export default BrowserPage;

