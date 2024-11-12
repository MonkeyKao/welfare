/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    extend: {
      fontSize:{
        'H1':'36px', //重點強調文字
        'H2':'24px', //頁面標題
        'H3':'16px', //區塊標題、功能卡片標題、試算結果
        'H4':'14px', //內文主要文字 ( 列表、彈窗、錯訊 )
        'H5':'12px', //盡量不使用
      },
    },
  },
  plugins: [],
}

