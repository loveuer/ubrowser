<template>
  <div class="title">
    <div class="icon">
      <svg  viewBox="0 0 1024 1024" version="1.1" width="30" height="30"><path d="M873.472 867.328H150.528l-40.96-40.96V197.632l40.96-40.96h178.176l38.912 28.672 34.816 108.544h471.04l40.96 40.96v492.544l-40.96 39.936z m-681.984-81.92h642.048V374.784h-460.8l-38.912-28.672-34.816-108.544H191.488v547.84z" fill="#437DFF" p-id="1224"></path><path d="M837.632 272.384h-61.44v-40.96h-358.4v-61.44h389.12l30.72 30.72z" fill="#63F7DE" p-id="1225"></path></svg> 
    </div>
    <div class="search">
      <div>
        <div class="s-icon">
          <svg viewBox="0 0 1024 1024" version="1.1" width="20" height="20"><path d="M1024 953.6L953.6 1024l-217.6-217.6 70.4-70.4 217.6 217.6z m-614.4-236.8c-166.4 0-307.2-140.8-307.2-307.2 0-166.4 140.8-307.2 307.2-307.2 166.4 0 307.2 140.8 307.2 307.2 0 166.4-140.8 307.2-307.2 307.2zM409.6 0C185.6 0 0 185.6 0 409.6s185.6 409.6 409.6 409.6c224 0 409.6-185.6 409.6-409.6S633.6 0 409.6 0z" fill="#D8D8D8" p-id="1999"></path></svg> 
        </div>
        <input type="text" name="search" id="search" placeholder="搜索..." />
      </div> 
    </div>
    <div class="menus"></div>
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
  height: 50px;
  align-items: center;
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
}
.title > .search > div {
  background: #f5f5f5;
  border-radius: 3px;
  display: flex;
  align-items: center;
  min-width: 200px;
}
.title > .search > div > .s-icon {
  margin-right: 10px;
  margin-left: 10px;
}
.title > .search > div > input {
  border: 0;
  background: #f5f5f5;
  outline: none;
  margin: 0;
  padding: 0;
  height: 40px;
  font-size: 1rem;
}
</style>
