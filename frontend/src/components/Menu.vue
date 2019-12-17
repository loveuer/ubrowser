<template>
  <div class="title">
    <div class="icon">
      <svg  viewBox="0 0 1024 1024" version="1.1" width="30" height="30"><path d="M873.472 867.328H150.528l-40.96-40.96V197.632l40.96-40.96h178.176l38.912 28.672 34.816 108.544h471.04l40.96 40.96v492.544l-40.96 39.936z m-681.984-81.92h642.048V374.784h-460.8l-38.912-28.672-34.816-108.544H191.488v547.84z" fill="#437DFF" p-id="1224"></path><path d="M837.632 272.384h-61.44v-40.96h-358.4v-61.44h389.12l30.72 30.72z" fill="#63F7DE" p-id="1225"></path></svg> 
    </div>
    <div class="search">
      <div>
        <div class="s-icon">
          <svg viewBox="0 0 1024 1024" version="1.1" width="15" height="15"><path d="M1024 953.6L953.6 1024l-217.6-217.6 70.4-70.4 217.6 217.6z m-614.4-236.8c-166.4 0-307.2-140.8-307.2-307.2 0-166.4 140.8-307.2 307.2-307.2 166.4 0 307.2 140.8 307.2 307.2 0 166.4-140.8 307.2-307.2 307.2zM409.6 0C185.6 0 0 185.6 0 409.6s185.6 409.6 409.6 409.6c224 0 409.6-185.6 409.6-409.6S633.6 0 409.6 0z" fill="#546e7a" p-id="1999"></path></svg> 
        </div>
        <input type="text" name="search" id="search" placeholder="搜索..." />
      </div> 
    </div>
    <div class="menus">
       <div class="m-download">
        <svg viewBox="0 0 1024 1024" version="1.1" width="25" height="25"><path d="M213.333333 853.333333h597.333334v-85.333333H213.333333m597.333334-384h-170.666667V128H384v256H213.333333l298.666667 298.666667 298.666667-298.666667z" fill="#546e7a" p-id="20511"></path></svg>
       </div>
       <div class="m-upload">
        <svg viewBox="0 0 1024 1024" version="1.1" width="25" height="25"><path d="M384 682.666667v-256H213.333333l298.666667-298.666667 298.666667 298.666667h-170.666667v256H384m-170.666667 170.666666v-85.333333h597.333334v85.333333H213.333333z" fill="#546e7a" p-id="20632"></path></svg> 
       </div>
       <div class="m-info">
      <svg viewBox="0 0 1024 1024" version="1.1" width="25" height="25"><path d="M554.666667 384h-85.333334V298.666667h85.333334m0 426.666666h-85.333334v-256h85.333334m-42.666667-384A426.666667 426.666667 0 0 0 85.333333 512a426.666667 426.666667 0 0 0 426.666667 426.666667 426.666667 426.666667 0 0 0 426.666667-426.666667A426.666667 426.666667 0 0 0 512 85.333333z" fill="#546e7a" p-id="20511"></path></svg> 
       </div>
       <div class="m-selector">
       <svg viewBox="0 0 1024 1024" version="1.1" width="25" height="25"><path d="M426.666667 725.333333l-213.333334-213.333333 60.16-60.586667L426.666667 604.586667l323.84-323.84L810.666667 341.333333m-298.666667-256A426.666667 426.666667 0 0 0 85.333333 512a426.666667 426.666667 0 0 0 426.666667 426.666667 426.666667 426.666667 0 0 0 426.666667-426.666667A426.666667 426.666667 0 0 0 512 85.333333z" fill="#546e7a" p-id="20632"></path></svg>
       </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Menu',
}
</script>

<style scoped>
.title{
  display: flex;
  height: 61px;
  top: 0px;
  position: fixed;
  left: 0px;
  width: 100%;
  align-items: center;
  margin-top: 5px;
  box-shadow: rgba(0, 0, 0, 0.1) 0 0 5px;
  border-bottom: 1px solid rgba(0,0,0,0.075);
  z-index: 1000;
  outline: none;
  margin-top: -1px;
}
.title > .icon {
  margin-left: 10px;
  border: 2px solid #54a0ff;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 5px;
  cursor: pointer;
}
.title > .search {
  margin-left: 10px;
  width: 400px;
  min-width: 235px;
}
.title > .search > div {
  background: #f5f5f5;
  border-radius: 5px;
  display: flex;
  align-items: center;
  min-width: 200px;
  padding-right: 5px;
}
.title > .search > div > .s-icon {
  margin-right: 10px;
  margin-left: 10px;
}
.title > .search > div > input {
  width: 100%;
  border: 0;
  background: #f5f5f5;
  outline: none;
  margin: 0;
  padding: 0;
  height: 40px;
  font-size: 1rem;
}

.title > .menus {
  display: flex;
  align-items: center;
  margin-left: auto;
  margin-right: 20px;
}
.title > .menus > div {
  margin-left: 15px;
  cursor: pointer;
  height: 40px;
  width: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.title > .menus > div:hover {
  background: #ddd;
  border-radius: 50%;
}
</style>
