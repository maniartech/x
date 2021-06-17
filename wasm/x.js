;(function (f) {
  // From: https://raw.githubusercontent.com/chilts/umd-template/master/template.js

  // module name and requires
  var name = 'x';
  var requires = [];

  // CommonJS
  if (typeof exports === "object" && typeof module !== "undefined") {
    module.exports = f.apply(null, requires.map(function(r) { return require(r); }));

  // RequireJS
  } else if (typeof define === "function" && define.amd) {
    define(requires, f);

  // <script>
  } else {
    var g;
    if (typeof window !== "undefined") {
      g = window;
    } else if (typeof global !== "undefined") {
      g = global;
    } else if (typeof self !== "undefined") {
      g = self;
    } else {
      // works providing we're not in "use strict";
      // needed for Java 8 Nashorn
      // seee https://github.com/facebook/react/issues/3037
      g = this;
    }

    g[name] = f.apply(null, requires.map(function(r) { return g[r]; }));
  }

})(function () {

  class XEvaluator {

    // Private constructor!
    constructor () {
      this.load("x.wasm")
    }

    /**
     * Loads the x.wasm from the specified URL.
     * @returns Returns true if the load is successful otherwise false.
     */
    load = (wsUri) => {

      // Check if x.wasm already loaded!
      if (typeof window.goEval === "function") {
        return true
      }

      // Load it
      const go = new Go();
      WebAssembly.instantiateStreaming(fetch(wsUri), go.importObject).then((result) => {
          go.run(result.instance);
      });

      // Check and return true if the wasm file has successfully loaded.
      return window.goEval !== undefined
    }

    /**
     * Evalueates the expression and return the results. Throws and EvalError
     * when an error is occurred while processing the expression.
     *
     * @param {string} expr The expression which needs to be evaluated
     * @param {object} context The context object, containts additional variables
     * @returns The evaluation results!
     *
     * @example
     * const result = x.eval("x + y", { x: 10, y: 20}) // Returns 30
     */
    eval = (expr, context) => {
      if (window.goEval === undefined) {
        throw new Error("The evaluator not available. Ensure that x.wasm is loaded!")
      }

      let result = window.goEval(expr, JSON.stringify(context))
      let strResult = result.toString()
      const errPrefix = "xError:"
      if (strResult.startsWith(errPrefix)) {
        throw new EvalError(strResult.substring(errPrefix.length))
      }
      return result
    }
  }

  return new XEvaluator()

});

