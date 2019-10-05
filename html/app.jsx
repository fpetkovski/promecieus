class SearchBar extends React.Component {
  constructor(props) {
    super(props);

    this.handleInputChange = this.handleInputChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleInputChange(event) {
    this.props.onSearchInput(event.target.value);
  }

  handleSubmit(event) {
    this.props.onSearchSubmit(event);
  }

  render() {
    return (
      <ReactBootstrap.Form horizontal>
        <ReactBootstrap.FormGroup>
          <ReactBootstrap.Row>
            <ReactBootstrap.Col xs={11}>
              <ReactBootstrap.FormControl
                autoFocus="true"
                type="text"
                placeholder="Feed me Prow URLs..."
                value={this.props.searchInput}
                onChange={this.handleInputChange}
              />
            </ReactBootstrap.Col>
            <ReactBootstrap.Col xs={1}>
              <ReactBootstrap.Button
                type="submit"
                onClick={this.handleSubmit}>
                Generate
              </ReactBootstrap.Button>
            </ReactBootstrap.Col>
          </ReactBootstrap.Row>
        </ReactBootstrap.FormGroup>
      </ReactBootstrap.Form>
    );
  }
}

class DeleteAppButton extends React.Component {
  render() {
    return (
      <ReactBootstrap.Button variant="warning" onClick={this.props.onDeleteApp}>
      Delete {this.props.appName} instance
      </ReactBootstrap.Button>
    )
  }
}

class Message extends React.Component {
  render() {
    var lines = this.props.message.split('/n');
    var variants = {
      "status": "info",
      "failure": "danger",
      "done": "success"
    }
    switch (this.props.action) {
      case 'done':
      case 'failure':
      case 'status':
        return (
          lines.map(item =>
              <ReactBootstrap.Alert variant={variants[this.props.action]}>
              {item}
              </ReactBootstrap.Alert>
          )
        )
        break;
      case 'link':
        return (
          <ReactBootstrap.Alert variant="warning">
            <ReactBootstrap.Alert.Link href={this.props.message}>{this.props.message}</ReactBootstrap.Alert.Link>
          </ReactBootstrap.Alert>
        )
        break;
      case 'app-label':
        return (
          <ReactBootstrap.Alert variant="warning">
            <DeleteAppButton onDeleteApp={this.props.onDeleteApp} appName={this.props.message}/>
          </ReactBootstrap.Alert>
        )
        break;
    }
  }
}

class Status extends React.Component {
  render() {
    if (this.props.messages.length == 0) {
      return (
        <span></span>
      );
    } else {
      return(
        <div>
          {
            this.props.messages.map(item =>
              <Message
                action={item.action}
                message={item.message}
                onDeleteApp={this.props.onDeleteApp}
              />
            )
          }
        </div>
      );
    }
  }
}

class SearchForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      searchInput: '',
      messages: [],
      appName: '',
      ws: null
    };

    this.handleSearchInput = this.handleSearchInput.bind(this);
    this.handleSearchSubmit = this.handleSearchSubmit.bind(this);
    this.handleDeleteApp = this.handleDeleteApp.bind(this);
    this.addMessage = this.addMessage.bind(this);
  }

  handleSearchInput(searchInput) {
    this.setState({searchInput: searchInput});
  }

  handleSearchSubmit(event) {
    event.preventDefault();
    if (this.state.searchInput.length == 0) {
      return;
    }

    try {
      this.state.ws.send(JSON.stringify({
        'action': 'new',
        'message': this.state.searchInput
      }))
    } catch (error) {
      console.log(error)
    }
  }

  handleDeleteApp(event) {
    console.log("handleDeleteApp " )
    try {
      this.state.ws.send(JSON.stringify({
        'action': 'delete',
        'message': this.state.appName
      }))
    } catch (error) {
      console.log(error)
    }
  }

  addMessage(message) {
    this.setState(state => ({ messages: [...state.messages, message] }))
    if (message.action === "app-label") {
      this.setState(state => ({appName: message.message}))
    }
  }

  componentDidMount() {
    var loc = window.location, new_uri;
    var ws_uri;
    if (loc.protocol === "https:") {
        ws_uri = "wss:";
    } else {
        ws_uri = "ws:";
    }
    ws_uri += "//" + loc.host;
    ws_uri += "/ws/status";
    var ws = new WebSocket(ws_uri);
    let that = this;
    var connectInterval;

    // websocket onopen event listener
    ws.onopen = () => {
      console.log("websocket connected");

      this.setState({ ws: ws });

      that.timeout = 250; // reset timer to 250 on open of websocket connection
      clearTimeout(connectInterval); // clear Interval on on open of websocket connection
    };

    // websocket onclose event listener
    ws.onclose = e => {
      console.log(
        `Socket is closed. Reconnect will be attempted in ${Math.min(
            10000 / 1000,
            (that.timeout + that.timeout) / 1000
        )} second.`,
        e.reason
      );

      that.timeout = that.timeout + that.timeout; //increment retry interval
      connectInterval = setTimeout(this.check, Math.min(10000, that.timeout)); //call check function after timeout
    };

    // websocket onerror event listener
    ws.onerror = err => {
      console.error(
        "Socket encountered error: ",
        err.message,
        "Closing socket"
      );

      ws.close();
    };

    ws.onmessage = evt => {
      console.log(evt.data);
      const message = JSON.parse(evt.data);
      this.addMessage(message)
    }
  }

  render() {
    let messages;
    let searchClass;
    if(this.state.appName != '') {
        messages =
        <Status
          messages={this.state.messages}
          onDeleteApp={this.handleDeleteApp}
        />
        searchClass = null
    } else {
        messages = []
        searchClass = 'search-center'
    }
    return (
      <div className={searchClass}>
        <h3>PromeCIeus</h3>
        <SearchBar
          searchInput={this.state.searchInput}
          onSearchInput={this.handleSearchInput}
          onSearchSubmit={this.handleSearchSubmit}
        />
        <br />
        {messages}
      </div>
    );
  }
}

ReactDOM.render(
  <SearchForm />,
  document.getElementById('container')
);
