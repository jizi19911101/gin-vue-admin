function Person() {
    // Person() 构造函数定义 `this`作为它自己的实例.
    this.age = 0;
    console.log(this,"1111")
  
    setInterval(function growUp() {
      // 在非严格模式, growUp()函数定义 `this`作为全局对象,
      // 与在 Person()构造函数中定义的 `this`并不相同.
      this.age++;
      console.log(this.age,"2222")
    }, 1000);
  }
  
  var p = new Person();