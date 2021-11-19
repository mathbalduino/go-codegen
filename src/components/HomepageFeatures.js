import React from 'react';
import clsx from 'clsx';
import styles from './HomepageFeatures.module.css';
import CodeBlock from '@theme/CodeBlock';
import Link from '@docusaurus/Link';

const parseGoStructs = `package main

import (
  "fmt"
  "github.com/mathbalduino/go-codegen"
  "go/types"
)

func main() {
  parser, e := parser.NewGoParser("#your_pattern_string#", parser.Config{})
  if e != nil {
    panic(e)
  }
  
  // Will print the name of every struct inside the parsed code
  parser.IterateStructs(func(struct_ *types.TypeName, logger parser.LoggerCLI) error {
    fmt.Println(struct_.Name())
    return nil
  })
}
`

const parseGoInterfaces = `package main

import (
  "fmt"
  "github.com/mathbalduino/go-codegen"
  "go/types"
)

func main() {
  parser, e := parser.NewGoParser("#your_pattern_string#", parser.Config{})
  if e != nil {
    panic(e)
  }
  
  // Will print the name of every interface inside the parsed code
  parser.IterateInterfaces(func(interface_ *types.TypeName, logger parser.LoggerCLI) error {
    fmt.Println(interface_.Name())
    return nil
  })
}
`

const focusCode = `package main

import (
  "fmt"
  "github.com/mathbalduino/go-codegen"
)

func main() {
  // Will skip any interface that doesn't have its name
  // equal to "SomeInterface"
  cfg := parser.Config{
    Focus: &parser.FocusTypeName("SomeInterface")
  }
  
  parser, e := parser.NewGoParser("#your_pattern_string#", cfg)
  if e != nil {
    panic(e)
  }
  
  parser.IterateInterfaces(func(interface_ *types.TypeName, logger parser.LoggerCLI) error {
    // Callback will be executed only if the interface name
    // is equal to "SomeInterface"
    fmt.Println(interface_.Name())
    
    return nil
  })
}

`

const filesCode = `package main

import (
  "fmt"
  "github.com/mathbalduino/go-codegen"
  "github.com/mathbalduino/go-codegen/goFile"
)

func main() {
  parser, e := parser.NewGoParser("#your_pattern_string#", parser.Config{})
  if e != nil {
    panic(e)
  }
  // Will create the "exampleFile" inside the "#some_pkg_name#" package, using the "#the_some_pkg_path#" package import path
  file := goFile.New("exampleFile", "#some_pkg_name#", "#the_some_pkg_path#")
  parser.IterateInterfaces(func(interface_ *types.TypeName, logger parser.LoggerCLI) error {
    // Just an example of an extremely simplified code generation. The generated
    // file will have a constant string with the name of every parsed interface
    file.AddCode(fmt.Sprintf("const InterfaceName_%s = \\"%s\\"\\n", interface_.Name(), interface_.Name()))
    return nil
  })
  // The file will be saved to the "#some_folder_path" folder path
  // The "#some_title#" will be written to the file header (comment section)
  e = file.Save("#some_title#", "#some_folder_path")
  if e != nil {
    panic(e)
  }
}
`

const FeatureList = [
  {
    title: 'Parse GO Structs',
    header: <CodeBlock className="language-go">{parseGoStructs}</CodeBlock>,
    description: (
      <>
        Create a new instance of the GO parser and you will be able to parse
        and gather information about the structs inside your code, using the <Link to={'/docs/go-parser-api#iterate-structs'}>IterateStructs</Link>
      </>
    ),
  },
  {
    title: 'Parse GO Interfaces',
    header: <CodeBlock className="language-go">{parseGoInterfaces}</CodeBlock>,
    description: (
      <>
        You can gather information about the interfaces too, just use the <Link to={'/docs/go-parser-api#iterate-interfaces'}>IterateInterfaces</Link>
      </>
    ),
  },
  {
    title: 'Focus',
    header: <CodeBlock className="language-go">{focusCode}</CodeBlock>,
    description: (
      <>
        If you don't want to iterate over all the parsed code, you can use
        the <Link to={'/docs/go-parser-api#focus'}>Focus</Link> feature to
        skip files, packages or typenames
      </>
    ),
  },
  {
    title: 'Files abstraction',
    header: (
      <CodeBlock className="language-go">{filesCode}</CodeBlock>
    ),
    description: (
      <>
        The library comes with builtin support for <Link to={'/docs/go-file'}>GO</Link> and <Link to={'/docs/ts-file'}>TS</Link> files, including
        formatting, import list handling (trust me, you don't want to handle it) and persistence
      </>
    ),
  },
];

function Feature({header, title, description}) {
  return (
    <div className={clsx('col col--6')}>
      {header}
      <div className="text--center padding-horiz--md">
        <h3>{title}</h3>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
