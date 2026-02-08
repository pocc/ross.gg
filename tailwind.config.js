/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './layouts/**/*.html',
    './content/**/*.md',
  ],
  darkMode: 'class',
  theme: {
    fontFamily: {
      display: ['"Fraunces"', 'Georgia', 'serif'],
      body: ['"Inter"', 'system-ui', 'sans-serif'],
      mono: ['"JetBrains Mono"', '"Fira Code"', 'monospace'],
    },
    fontSize: {
      xs:   ['0.72rem',  { lineHeight: '1.4' }],
      sm:   ['0.889rem', { lineHeight: '1.5' }],
      base: ['1.125rem', { lineHeight: '1.7' }],
      lg:   ['1.25rem',  { lineHeight: '1.6' }],
      xl:   ['1.563rem', { lineHeight: '1.4' }],
      '2xl':  ['1.953rem', { lineHeight: '1.3' }],
      '3xl':  ['2.441rem', { lineHeight: '1.2' }],
      '4xl':  ['3.052rem', { lineHeight: '1.1' }],
      '5xl':  ['3.815rem', { lineHeight: '1.05' }],
    },
    extend: {
      colors: {
        noir: {
          950: '#0a0a0b',
          900: '#111113',
          850: '#18181b',
          800: '#1f1f23',
          700: '#2a2a30',
          600: '#3a3a42',
        },
        stone: {
          50:  '#fafaf9',
          100: '#f5f5f4',
          200: '#e7e5e4',
          300: '#d6d3d1',
          400: '#a8a29e',
          500: '#78716c',
        },
        fg: {
          DEFAULT: 'var(--fg)',
          muted: 'var(--fg-muted)',
          faint: 'var(--fg-faint)',
        },
        accent: {
          DEFAULT: 'var(--accent)',
          light: '#f0c078',
          dark: '#c08030',
          muted: 'rgba(224, 164, 88, 0.2)',
        },
        seedling: '#6ec46e',
        budding: '#e0a458',
        evergreen: '#4a90d9',
      },
      maxWidth: {
        prose: '68ch',
        wide: '72rem',
      },
      borderRadius: {
        card: '1rem',
      },
      boxShadow: {
        card: '0 1px 3px rgba(0,0,0,0.08), 0 1px 2px rgba(0,0,0,0.06)',
        'card-hover': '0 8px 24px rgba(0,0,0,0.12)',
      },
      typography: ({ theme }) => ({
        DEFAULT: {
          css: {
            '--tw-prose-body': 'var(--fg)',
            '--tw-prose-headings': 'var(--fg)',
            '--tw-prose-links': 'var(--accent)',
            '--tw-prose-bold': 'var(--fg)',
            '--tw-prose-code': 'var(--accent)',
            '--tw-prose-pre-bg': '#1f1f23',
            '--tw-prose-pre-code': '#e8e6e3',
            fontFamily: theme('fontFamily.body').join(', '),
            maxWidth: '68ch',
            h1: { fontFamily: theme('fontFamily.display').join(', '), fontWeight: '700', letterSpacing: '-0.02em' },
            h2: { fontFamily: theme('fontFamily.display').join(', '), fontWeight: '700', letterSpacing: '-0.01em' },
            h3: { fontFamily: theme('fontFamily.display').join(', '), fontWeight: '700' },
            a: {
              color: 'var(--accent)',
              textDecoration: 'underline',
              textDecorationColor: 'rgba(224, 164, 88, 0.3)',
              textUnderlineOffset: '3px',
              '&:hover': { textDecorationColor: 'var(--accent)' },
            },
            code: { fontFamily: theme('fontFamily.mono').join(', '), fontWeight: '400' },
            'code::before': { content: 'none' },
            'code::after': { content: 'none' },
          },
        },
      }),
    },
  },
  plugins: [
    require('@tailwindcss/typography'),
  ],
};
