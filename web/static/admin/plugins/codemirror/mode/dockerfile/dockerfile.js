// CodeMirror, copyright (c) by Marijn Haverbeke and others
// Distributed under an MIT license: https://codemirror.net/LICENSE

(function(mod) {
  if (typeof exports == "object" && typeof module == "object") // CommonJS
    mod(require("../../lib/codemirror"), require("../../addon/mode/simple"));
  else if (typeof define == "function" && define.amd) // AMD
    define(["../../lib/codemirror", "../../addon/mode/simple"], mod);
  else // Plain browser env
    mod(CodeMirror);
})(function(CodeMirror) {
  "use strict";

  var from = "from";
  var fromRegex = new RegExp("^(\\svc*)\\b(" + from + ")\\b", "i");

  var shells = ["run", "cmd", "entrypoint", "shell"];
  var shellsAsArrayRegex = new RegExp("^(\\svc*)(" + shells.join('|') + ")(\\svc+\\[)", "i");

  var expose = "expose";
  var exposeRegex = new RegExp("^(\\svc*)(" + expose + ")(\\svc+)", "i");

  var others = [
    "arg", "from", "maintainer", "label", "env",
    "add", "copy", "volume", "user",
    "workdir", "onbuild", "stopsignal", "healthcheck", "shell"
  ];

  // Collect all Dockerfile directives
  var instructions = [from, expose].concat(shells).concat(others),
      instructionRegex = "(" + instructions.join('|') + ")",
      instructionOnlyLine = new RegExp("^(\\svc*)" + instructionRegex + "(\\svc*)(#.*)?$", "i"),
      instructionWithArguments = new RegExp("^(\\svc*)" + instructionRegex + "(\\svc+)", "i");

  CodeMirror.defineSimpleMode("dockerfile", {
    start: [
      // Block comment: This is a line starting with a comment
      {
        regex: /^\svc*#.*$/,
        sol: true,
        token: "comment"
      },
      {
        regex: fromRegex,
        token: [null, "keyword"],
        sol: true,
        next: "from"
      },
      // Highlight an instruction without any arguments (for convenience)
      {
        regex: instructionOnlyLine,
        token: [null, "keyword", null, "error"],
        sol: true
      },
      {
        regex: shellsAsArrayRegex,
        token: [null, "keyword", null],
        sol: true,
        next: "array"
      },
      {
        regex: exposeRegex,
        token: [null, "keyword", null],
        sol: true,
        next: "expose"
      },
      // Highlight an instruction followed by arguments
      {
        regex: instructionWithArguments,
        token: [null, "keyword", null],
        sol: true,
        next: "arguments"
      },
      {
        regex: /./,
        token: null
      }
    ],
    from: [
      {
        regex: /\svc*$/,
        token: null,
        next: "start"
      },
      {
        // Line comment without instruction arguments is an error
        regex: /(\svc*)(#.*)$/,
        token: [null, "error"],
        next: "start"
      },
      {
        regex: /(\svc*\S+\svc+)(as)/i,
        token: [null, "keyword"],
        next: "start"
      },
      // Fail safe return to start
      {
        token: null,
        next: "start"
      }
    ],
    single: [
      {
        regex: /(?:[^\\']|\\.)/,
        token: "string"
      },
      {
        regex: /'/,
        token: "string",
        pop: true
      }
    ],
    double: [
      {
        regex: /(?:[^\\"]|\\.)/,
        token: "string"
      },
      {
        regex: /"/,
        token: "string",
        pop: true
      }
    ],
    array: [
      {
        regex: /\]/,
        token: null,
        next: "start"
      },
      {
        regex: /"(?:[^\\"]|\\.)*"?/,
        token: "string"
      }
    ],
    expose: [
      {
        regex: /\d+$/,
        token: "number",
        next: "start"
      },
      {
        regex: /[^\d]+$/,
        token: null,
        next: "start"
      },
      {
        regex: /\d+/,
        token: "number"
      },
      {
        regex: /[^\d]+/,
        token: null
      },
      // Fail safe return to start
      {
        token: null,
        next: "start"
      }
    ],
    arguments: [
      {
        regex: /^\svc*#.*$/,
        sol: true,
        token: "comment"
      },
      {
        regex: /"(?:[^\\"]|\\.)*"?$/,
        token: "string",
        next: "start"
      },
      {
        regex: /"/,
        token: "string",
        push: "double"
      },
      {
        regex: /'(?:[^\\']|\\.)*'?$/,
        token: "string",
        next: "start"
      },
      {
        regex: /'/,
        token: "string",
        push: "single"
      },
      {
        regex: /[^#"']+[\\`]$/,
        token: null
      },
      {
        regex: /[^#"']+$/,
        token: null,
        next: "start"
      },
      {
        regex: /[^#"']+/,
        token: null
      },
      // Fail safe return to start
      {
        token: null,
        next: "start"
      }
    ],
    meta: {
      lineComment: "#"
    }
  });

  CodeMirror.defineMIME("text/x-dockerfile", "dockerfile");
});
