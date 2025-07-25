/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./components/**/*.{js,vue,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    "./nuxt.config.{js,ts}",
    "./app.vue"
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          50: '#f0fdf4',
          100: '#dcfce7',
          200: '#bbf7d0',
          300: '#86efac',
          400: '#4ade80',
          500: '#22c55e',
          600: '#16a34a',
          700: '#15803d',
          800: '#166534',
          900: '#14532d',
          950: '#052e16',
        },
        green: {
          50: '#f0fff4',
          100: '#dcfce7',
          200: '#bbf7d0',
          300: '#86efac',
          400: '#00ff41',
          500: '#00cc33',
          600: '#16a34a',
          700: '#15803d',
          800: '#166534',
          900: '#003d0f',
          950: '#052e16',
        },
        gray: {
          50: '#f9fafb',
          100: '#f3f4f6',
          200: '#e5e7eb',
          300: '#d1d5db',
          400: '#9ca3af',
          500: '#6b7280',
          600: '#4b5563',
          700: '#374151',
          800: '#1f2937',
          900: '#0a0a0a',
          950: '#050505',
        }
      },
      fontFamily: {
        mono: ['Courier New', 'monospace'],
        sans: ['IBM Plex Sans SC', 'PingFang SC', 'Microsoft YaHei', 'sans-serif'],
        chinese: ['IBM Plex Sans SC', 'PingFang SC', 'Microsoft YaHei', 'sans-serif'],
      },
      animation: {
        'float': 'float 3s ease-in-out infinite',
        'pulse': 'pulse 2s ease-in-out infinite',
        'data-flow': 'dataFlow 2s linear infinite',
        'rotate-3d': 'rotate3d 10s linear infinite',
        'slide-in-left': 'slideInLeft 0.3s ease-out',
      },
      keyframes: {
        float: {
          '0%, 100%': { transform: 'translateY(0px)' },
          '50%': { transform: 'translateY(-10px)' },
        },
        dataFlow: {
          '0%': { transform: 'translateX(-100%)', opacity: '0' },
          '50%': { opacity: '1' },
          '100%': { transform: 'translateX(100%)', opacity: '0' },
        },
        rotate3d: {
          '0%': { transform: 'rotateX(0deg) rotateY(0deg)' },
          '100%': { transform: 'rotateX(360deg) rotateY(360deg)' },
        },
        slideInLeft: {
          'from': { opacity: '0', transform: 'translateX(-20px)' },
          'to': { opacity: '1', transform: 'translateX(0)' },
        },
      },
      boxShadow: {
        'glow': '0 0 20px rgba(0, 255, 65, 0.3)',
        'glow-lg': '0 0 30px rgba(0, 255, 65, 0.4)',
      },
      backgroundImage: {
        'grid-pattern': 'linear-gradient(rgba(0, 255, 65, 0.1) 1px, transparent 1px), linear-gradient(90deg, rgba(0, 255, 65, 0.1) 1px, transparent 1px)',
      },
      backgroundSize: {
        'grid': '50px 50px',
      },
    },
  },
  plugins: [
    function({ addUtilities }) {
      const newUtilities = {
        '.glow': {
          'box-shadow': '0 0 20px rgba(0, 255, 65, 0.3)',
        },
        '.text-glow': {
          'text-shadow': '0 0 10px #00ff41',
        },
        '.float-animation': {
          'animation': 'float 3s ease-in-out infinite',
        },
        '.pulse-animation': {
          'animation': 'pulse 2s ease-in-out infinite',
        },
        '.data-flow': {
          'animation': 'dataFlow 2s linear infinite',
        },
      }
      addUtilities(newUtilities)
    }
  ],
}